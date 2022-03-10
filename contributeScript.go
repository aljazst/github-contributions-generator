package main

import (
    "flag"
    "fmt"
    "os"
    "log"
    "time"
    "strings"
    "strconv"
    "math/rand"
    "math"
    "os/exec"

)

const FILE_NAME string = "data.txt"
const DATE_FORMAT string = "2006-01-02 15:04:05"

const ROWS int = 7
const COLUMNS int = 5

var MESSAGE_OPACITY int = 2 // default contribution in the letter. 2x per day

func writeErrorMessage(err error) {
    if err != nil {
        log.Println("There has been an error: ",err)
        return
    }
}


func contribute(commit_date string) {
    // usefull links: https://golangbot.com/write-files/

    // file mode is in octal notation for the user, group and other (indicated by the 0, then we have 6 for 110, and 4 for 100. rwx)
    // https://stackoverflow.com/a/18415935/7973144 and https://ss64.com/bash/chmod.html and https://docs.nersc.gov/filesystems/unix-file-permissions/
    file, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Println("OpenFile err: ",err)
        return
    }
    defer file.Close()

    _, err2 := file.WriteString(commit_date+"\n\n")

    if err2 != nil {
        log.Fatal("WriteString err: ", err2)
    }
    //https://pkg.go.dev/os#File.Sync just in case we flush.
    file.Sync()
    //randomNum := rand.Intn(100 - 1) + 3
    exec.Command("git", "add", ".").Run()
    exec.Command("git", "commit", "-m", "Commit date was: "+ commit_date, "--date", commit_date).Run()

}

func generateDate(commit_limit, frequency int, time_period [2]int) {
    currentTime := time.Now()

    //https://pkg.go.dev/time#Time.AddDate -> AddDate(year,month,day)
    //startCommitDate := currentTime.AddDate(-1,0,0)
    var date string
    //rand.Intn(max - min) + min
    for i := 0; i < time_period[1]; i++ {
        rndNumOfCommits := rand.Intn(commit_limit - 1) + 1
        n := 0
        if rand.Intn(100 + 1) <= frequency {
            for n < rndNumOfCommits {
                date = currentTime.AddDate(-1,time_period[0],i).Format(DATE_FORMAT)
                n++
                contribute(date)

            }
        }
    }
}
func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
func getNumberOfDaysBetweenMonths(startMonth, endMonth int) int {
    currentYear, _, _ := time.Now().Date()
    return int(math.Round(Date(currentYear, endMonth, 0).Sub(Date(currentYear, startMonth, 0)).Hours() / 24))
}

func contributionsPerDay(num int) int {
    if num >= 15 {
        num = 15
    }
    if num < 2 {
        num = 2
    }
    return num
}

func contributins_specific_months(specified_months string) [2]int {
    //allright, this isn't the best solution but, its currently the only one I can come up with

    //an array that stores 2 values. The firs value defines the start month and the second is the number of days that the 2 month have between them
    //so when we read the [1] value we know how many times we have to contribute, untill we reach the desired time period, default walue is 365 days.
    months := [2]int {0,12}

    _, currentMonth, _ := time.Now().Date() //a hack to get the month type in int. The Month() function returns the name of the month in a string.

    if strings.Contains(specified_months,"-") {
        //if its the default value we just commit the entire year.
        if specified_months == "1-12" {
            months[0] = 0
            months[1] = getNumberOfDaysBetweenMonths(0, 12)
        } else {
           
        sliceOfMonths := strings.Split(specified_months, "-")

        startMonth, err := strconv.Atoi(sliceOfMonths[0])
        writeErrorMessage(err)

        endMonth, err1 := strconv.Atoi(sliceOfMonths[1])
        writeErrorMessage(err1)
        
        if startMonth - int(currentMonth) < 0 {
            months[0] = 0
        } else {
            months[0] = startMonth - int(currentMonth) //if we want to commit on specific months we have to subtract the month we are currently in.
        }
        
        months[1] = getNumberOfDaysBetweenMonths(startMonth, endMonth  + 1)
        fmt.Println("Number of days with possible commits:", months[1]) 
        }
    } else {
       startMonth, err := strconv.Atoi(specified_months)
       writeErrorMessage(err)
   
       if startMonth - int(currentMonth) < 0 {
            months[0] = 0
        } else {
            months[0] = startMonth - int(currentMonth) //if we want to commit on specific months we have to subtract the month we are currently in.
        }
        // get the number of days in one specific month. we use startMonth +1 because the functions calculates the number of days between 2 months
       months[1] = getNumberOfDaysBetweenMonths(startMonth,startMonth + 1)
       fmt.Println("Number of days with possible commits:", months[1]) 
    }

    //fmt.Println("the functions contribute_specific_months_returned: ", months)
    return months
}

func runScript(repository, timePeriod string, commit_limit, frequency int, message, backgroundSaturation string, dayCounter int, isRandom bool) {

    // TODO: error handling of the os commands
    if isRandom {
        os.MkdirAll("randomContributions", os.ModeDir)

        os.Chdir("randomContributions")
        os.RemoveAll(".git")
        os.RemoveAll("data.txt")
        exec.Command("git", "init").Run()
    
    } else {

       os.MkdirAll("nonrandomContributions", os.ModeDir)

       os.Chdir("nonrandomContributions")
       os.RemoveAll(".git")
       os.RemoveAll("data.txt")
       exec.Command("git", "init").Run()
    }



    if repository != ""{
        if isRandom {
            generateDate(contributionsPerDay(commit_limit), frequency, contributins_specific_months(timePeriod))
        } else {

            if backgroundSaturation != "not_defined" {

                generateDate(setSaturation(backgroundSaturation))
            }

            runNonRandomScript(message, dayCounter)

        }

        //TODO: error handling of the git commands -> https://pkg.go.dev/os/exec#Command
        exec.Command("git", "branch", "-M", "main").Run()
        exec.Command("git", "remote","add", "origin", repository).Run()

        exec.Command("git", "push", "-u", "origin", "main").Run()

    } else {
        fmt.Println("Holdup. You just wanted to run a github contrubutins script without entering a github repo? Try again.")
    }
    

    fmt.Println("")
    fmt.Println("We are done here! It may take some time for GitHub to display the changes. It depends on how many commits we have made.")
    fmt.Println("")

}

func setSaturation(saturationLevel string) (commit_level int, freq int, months [2]int) {

    if saturationLevel == "low" || saturationLevel == "Low" {
        MESSAGE_OPACITY = 5
        return 2, 100, [2]int {0,365}
    } else if saturationLevel == "medium" || saturationLevel == "Medium" {
        MESSAGE_OPACITY = 9
        return 4, 100, [2]int {0,365}
    } else if saturationLevel == "high" || saturationLevel == "High" {
        MESSAGE_OPACITY = 15
        return 7, 100, [2]int {0,365}
    } /*else if saturationLevel == "none" || saturationLevel == "None" {
        MESSAGE_OPACITY = 15
        return 0, 0, [2]int {1,12}
    } */

    //the low option will be the default
    return 2, 100, [2]int {1,12}
}

var letterMatrix [][]int 
var number int = 0
var iterationCounter int = 0
var letterCounter int = 0
var enteredMessage string 

func runNonRandomScript(message string, dayCounter int) {


    //now this is an ugly hack. Because the flag returns a pointer of the message the slice operations in this function change the value of the message. So we save the original value of the message and use it.
    if letterCounter == 0 {
        enteredMessage = message
    }
    
    number++
    messageLength := len(enteredMessage)
    fmt.Println("Message length: ",messageLength)
    fmt.Println("Message is: ", enteredMessage)
    
    var date string
    currentTime := time.Now()
    if iterationCounter < 1 {
        letterMatrix = ReturnMatrix(enteredMessage)
    }
    for i := 0; i < COLUMNS; i++ {
        for j := 0; j < ROWS; j++ {
   
            if letterMatrix[j][i] == 1 {
                 date = currentTime.AddDate(-1, 0, dayCounter).Format(DATE_FORMAT)
                  for k := 0; k < MESSAGE_OPACITY; k++ {         
                    contributeSpecific(date)
                  }
                }  
                
            dayCounter++
            iterationCounter++
            fmt.Println("Iteration: ", iterationCounter, "/", setLength(messageLength))

        }
    }

    if letterCounter < messageLength - 1 && letterCounter <= 10 {
        letterCounter++
    }

    if iterationCounter < setLength(messageLength) {

        letterMatrix = ReturnMatrix(enteredMessage[letterCounter:])
        fmt.Print(letterMatrix)
        
        //recursion. Call this function if there is another letter to write
        if number != iterationCounter {
        runNonRandomScript(enteredMessage[letterCounter:], dayCounter)
        }

    }

}

// one letter takes up 35 squares. So we have a 7x5 matrix. Depending on how many letters we have in the message, thata the number of squares we want to "paint"
func setLength(messagelength int) int {

    switch messagelength {
    case 1:
        return 35
    case 2:
        return 70
    case 3:
        return 105
    case 4:
        return 140
    case 5:
        return 175
    case 6:
        return 210
    case 7:
        return 245
    case 8:
        return 280
    case 9:
        return 315
    case 10:
        return 350
    }

    return 365
}

func contributeSpecific (date string) {

    file, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Println("OpenFile err: ",err)
        return
    }
    defer file.Close()

    _, err2 := file.WriteString(date+"\n\n")

    if err2 != nil {
        log.Fatal("WriteString err: ", err2)
    }
    //https://pkg.go.dev/os#File.Sync just in case we flush.
    file.Sync()

    exec.Command("git", "add", ".").Run()
    exec.Command("git", "commit", "-m", "Commit date was: "+ date, "--date", date).Run()


}

func main() {

    randomFlag := flag.NewFlagSet("random", flag.ExitOnError)

    repository := randomFlag.String("repository","","Enter a link to an empty non-initialized GitHub repository to which you want to push the generated file. The link can be an SSH (assuming you have an ssh key) or the HTTPS format. (e.g., git@github.com:yourusername/yourrepo.git or https://github.com/yourusername/yourrepo.git) ")
    commitLimmit := randomFlag.Int("commit_limit", 7, "Set the limit(n) of commits per single day. The script will randomly commit from 1 to n times a day. The maximum is 15 and the minimum is 1.")
    frequency := randomFlag.Int("frequency", 85, "The procentage of days out of 365 you would like to contribute. E.g., if you enter 20, you will contribute 73 days out of 1 year.")
    timePeriod := randomFlag.String("month","1-12", "Contribute only in a specific period. If you enter 3-5, you will only commit from march(starting the same day of month as today) to may. Entering only one number like 8(october) will prompt the script to commit only on the specified month.")


    
    
    nonrandomFlag := flag.NewFlagSet("nonrandom", flag.ExitOnError)
    repositoryNonRand := nonrandomFlag.String("repository","","Enter a link to an empty non-initialized GitHub repository to which you want to push the generated file. The link can be an SSH (assuming you have an ssh key) or the HTTPS format. (e.g., git@github.com:yourusername/yourrepo.git or https://github.com/yourusername/yourrepo.git) ")
    message := nonrandomFlag.String("message", "hello", "Enter the message you would like to be displayed on the contribution graph. The maximum ammount of characters is 10.")
    panelStartDate := nonrandomFlag.Int("adjust_date", 0, "Adjust the date difference between the current date and the first date displayed in the GitHub contributions panel. E.g, toda is the 8th of March, but GitHubs panel is still one day behind(7th of March). So you enter -1 to change todays date to the one displayed on the panel.")
    backgroundSaturation := nonrandomFlag.String("saturation", "not_defined", "The ammount of background saturation. Options are: low, medium and high. You should choose the option on behalf of your current contributions panel. If you have very little contributions choose the low option. If you have been very active, you will need to choose high in order for the text to be readable/visible. If you don't define the saturation, it won't be set. ")
   
    if len(os.Args) < 2 {
        fmt.Println("expected 'random' or 'nonrandom' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {

    case "random":
        randomFlag.Parse(os.Args[2:])
        fmt.Println("subcommand 'random'")
        fmt.Println("  repository:", *repository)
        fmt.Println("  commit limit:", *commitLimmit)
        fmt.Println("  time period (months):", *timePeriod)
        fmt.Println("  frequency of commits is:", *frequency, " % of the year")
        //fmt.Println("  random values you entered that are up to no good:", randomFlag.Args())
        runScript(*repository, *timePeriod, *commitLimmit, *frequency, *message, *backgroundSaturation ,*panelStartDate, true)
   
    case "nonrandom":
        nonrandomFlag.Parse(os.Args[2:])
        fmt.Println("subcommand 'nonrandom'")
        fmt.Println("  message:", *message)
        fmt.Println("  repository:", *repositoryNonRand)
        fmt.Println("  start date:", *panelStartDate)
        fmt.Println("  background saturation:", *backgroundSaturation)
        //fmt.Println("  tail:", nonrandomFlag.Args())

        runScript(*repositoryNonRand, *timePeriod, *commitLimmit, *frequency, *message, *backgroundSaturation, *panelStartDate, false)

    default:
        fmt.Println("Expected 'random' or 'nonrandom' subcommands!\n")
        fmt.Println("Arguments for random flag: \n")
        randomFlag.PrintDefaults()
        fmt.Println("\n\n")
        fmt.Println("Arguments for nonrandom flag: \n")
        nonrandomFlag.PrintDefaults()
        os.Exit(1)
    }
}