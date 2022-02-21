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
    randomNum := rand.Intn(100 - 1) + 3
    exec.Command("git", "add", ".").Run()
    exec.Command("git", "commit", "-m", string(randomNum), "--date", commit_date).Run()

}

func randate(commit_limit, frequency int, time_period [2]int) {
    currentTime := time.Now()
    fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))

    //https://pkg.go.dev/time#Time.AddDate -> AddDate(year,month,day)
    //startCommitDate := currentTime.AddDate(-1,0,0)
    var date string
    //rand.Intn(max - min) + min
    for i := 0; i < time_period[1]; i++ {
        rndNumOfCommits := rand.Intn(commit_limit - 1) + 1
        n := 0
        if rand.Intn(100 + 1) <= frequency {
            for n < rndNumOfCommits {
                date = currentTime.AddDate(-1,time_period[0],i).Format("2006-01-02 15:04:05")
                n++
                contribute(date)
            }
        }
    }

}
func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func contributins_specific_months(specified_months string) [2]int {
    //allright, this isn't the best solution but, its currently the only one I can come up with

    //an array that stores 2 values. The firs value defines the start month and the second is the number of days that the 2 month have between them
    //so when we read the [1] value we know how many times we have to contribute, untill we reach the desired time period
    var months [2]int
    _, month, _ := time.Now().Date() //a hack to get the month type in int. The Month() function returns the name of the month in a string.

    if strings.Contains(specified_months,"-") {
        //if its the default value we just commit the entire year.
        if specified_months == "1-12" {
            months[0] = 0
            months[1] = 12
        } else {
        sliceOfMonths := strings.Split(specified_months, "-")
        fmt.Println(sliceOfMonths)
        firstMonth := sliceOfMonths[0]
        secondMonth := sliceOfMonths[1]
        startMonth, err := strconv.Atoi(firstMonth)
        writeErrorMessage(err)

        teteg, err1 := strconv.Atoi(secondMonth)
        writeErrorMessage(err1)

        months[0] = startMonth - int(month) //if we want to commit on specific months we have to subtract the month we are currently in.
        
        days := Date(2006, teteg, 0).Sub(Date(2006, startMonth, 0)).Hours() / 24
        months[1] = int(math.Round(days))
        fmt.Println("number of days betweeen:", days) // 366
        }
    } else {
       startMonth, err := strconv.Atoi(specified_months)
       writeErrorMessage(err)
        //ttt :=time.Now()
        
       //month := time.Date(2006, startMonth, 0) this is going to give us the number od days in that specific month
       months[0] = startMonth - int(month)
       months[1] = 30 // for now we will use 30
    }

    fmt.Println(months)
    return months
}

func test(tmp, repository, timePeriod string, commit_limit, frequency int) {

    

    fmt.Println("the repo is", repository)
    fmt.Println("the commit limit is", commit_limit)
    fmt.Println("the time period is", timePeriod)
    /* if err := os.Mkdir("a", os.ModePerm); err != nil {
        log.Fatal(err)
    } */
   // out := exec.Command("mkdir", "lol").Run()
    out := os.MkdirAll("contributions", os.ModeDir)

    of := os.Chdir("contributions")
    os.RemoveAll(".git")
    os.RemoveAll("data.txt")
    exec.Command("git", "init").Run()

    if of != nil {
        log.Fatal(of)
    }
    
    if repository != ""{
        randate(commit_limit, frequency, contributins_specific_months(timePeriod))
    }
    


    exec.Command("git", "branch", "-M", "main").Run()
    exec.Command("git", "remote","add", "origin", repository).Run()

    exec.Command("git", "push", "-u", "origin", "main").Run()

    fmt.Println("done")

   // makeInit := exec.Command("rm", "-rf", ".git")

    //tmp1 := makeInit.Run()

    if out == nil  {
        fmt.Println("Command Successfully Executed")
        
    } else {
        fmt.Printf("%s", out)
        fmt.Printf("%s", out)
    }


}


func main() {

    randomFlag := flag.NewFlagSet("random", flag.ExitOnError)

    repository := randomFlag.String("repository","","Enter a link to an empty non-initialized GitHub repository to which you want to push the generated file. The link can be an SSH (assuming you have an ssh key) or the HTTPS format. (e.g., git@github.com:yourusername/yourrepo.git or https://github.com/yourusername/yourrepo.git) ")
    commitLimmit := randomFlag.Int("commit_limit", 7, "Set the limit(n) of commits per single day. The script will randomly commit from 1 to n times a day. The maximum is 15 and the minimum is 1.")
    frequency := randomFlag.Int("frequency", 85, "The procentage of days out of 365 you would like to contribute. E.g., if you enter 20, you will contribute 73 days out of 1 year.")
    timePeriod := randomFlag.String("month","1-12", "Contribute only in a specific period. If you enter 3-5, you will only commit from march(starting the same day of month as today) to may. Entering only one number like 8(october) will prompt the script to commit only on the specified month.")
    fooName := randomFlag.String("name", "", "name")


    nonrandomFlag := flag.NewFlagSet("nonrandom", flag.ExitOnError)
    barLevel := nonrandomFlag.Int("level", 0, "level")

    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {

    case "random":
        randomFlag.Parse(os.Args[2:])
        fmt.Println("subcommand 'random'")
        fmt.Println("  repository:", *repository)
        fmt.Println("  commit limit:", *commitLimmit)
        fmt.Println("  time period (months):", *timePeriod)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", randomFlag.Args())
        test(*fooName, *repository, *timePeriod, *commitLimmit, *frequency)
        //just call a function here and pass in the parameters
    case "nonrandom":
        nonrandomFlag.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", nonrandomFlag.Args())
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