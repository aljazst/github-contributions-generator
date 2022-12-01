# GitHub contributions generator

A script written in Go, which goal is to make your GitHub contributions graph the way you want it to be.
:shipit: The goal of this was to improve my Go skills and to have some fun! Enjoy! :beers:

## What's the goal? 
Let't transform this 
![original-github](https://user-images.githubusercontent.com/76880102/157703961-8ed845a2-4210-4b1c-b8fb-dbbeb3dd85ce.PNG)

to something more interesting

![bamboozled_med](https://user-images.githubusercontent.com/76880102/157708873-671f7c06-b5b2-490a-ad41-11173e450e2c.PNG)

## Disclaimer
This script was written for fun. In no way is it meant to "cheat". Nobody should judge a programmer regarding to his GitHub contributions graph.
It's still not perfect. If you will use the nonrandom (message) option, there can be some distorted letters.
**The creation of the directory part is implemented for Windows OS. If you use linux, check out the branch named patch-1.**

### System requirements
This script is written in Go, so you will need it. And ofcourse you will need git. 

## How to run the nonrandom option (displaying a message)
1. Create an fully empty GitHub repository (not even a README). Do not initialize it.
2. Example of a command:
```sh
go run . nonrandom -message="HelloWorld" -repository="https://github.com/yourUserName/yourRepo.git" -adjust_date=-3 -saturation=medium
```
Note the adjust_date parameter. It's a important one. More about it below. 

![helloworld_high](https://user-images.githubusercontent.com/76880102/157709681-304db8cb-e8b8-46b3-93a5-d37d4f40c698.PNG)

### The adjust_date parameter

If you want to display a message, you will have to check if the current date and the 1st date displayed in the GitHub contributions graph are the same. I have no idea how the graph works, but usually the first "box" (top left corner) is a few days behind the current date. 

![dateExxample](https://user-images.githubusercontent.com/76880102/157712749-0ab75ee0-01bf-4203-bc2d-979649414032.png)

In the example above we see that there is a 3 day difference between todays date and the first GitHub date. So we have to let the script know that it has to start 3 days earlier. So we use the -adjust_date=-3 flag. 

If you don't pay attention to this, the script will start contributing exactly 1 year ago from today, and this will happen:
![no_date_adjustment_example](https://user-images.githubusercontent.com/76880102/157713381-e4212f30-3826-4181-9f8c-6ba2cd7cdf5e.PNG)


### The saturation parameter (filling the background when displaying a message)
If you don't specify a background saturation level, the script will just write the message with no background. That can be a problem if you have commits from before as the squares will "collide" with the letters and make them unreadable. 
Example with no saturation parameter provided:
```sh
go run . nonrandom -message="sad->" -repository="https://github.com/yourUserName/yourRepo.git" -adjust_date=-3
```
![no_background](https://user-images.githubusercontent.com/76880102/157715626-94a741d9-452c-42d8-9d26-47b03c79418f.PNG)

#### low, medium, high background saturation. Which one to use?
It all depends on your GitHub contributions graph. 
If you have been very active, and made a lot of commits you will have to use the high option, you will fill out the backround to a level that will make your message readable. 

If you don't have a lot of commits you can go with the low or medium option. Take in to account that the high saturation option takes the most time to complete.

This is an example of a medium saturated background:
![barcelona_github](https://user-images.githubusercontent.com/76880102/157718345-0f90bf2b-88e8-4730-bc0d-23563df2205a.PNG)

## Paremeters
| Flag Name        | Type    | Required | Default                                        | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
|-------------------|---------|----------|------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| repository        | String  | true     |                                       | Enter a link to an empty non-initialized GitHub repository to which you want to push the generated file. The link can be an SSH (assuming you have an ssh key) or the HTTPS format. (e.g., git@github.com:yourusername/yourrepo.git or https://github.com/yourusername/yourrepo.git)                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| saturation          | String  | false     | none                                        | The ammount of background saturation. Options are: <b>low, medium and high</b>. You should choose the option on behalf of your current GitHub contributions panel. If you have very little contributions choose the low option. If you have been very active, you will need to choose high in order for the text to be readable/visible. If you don't define the saturation, it won't be set.                                                                                                                                                                                                                                                                                                       |
| adjust_date           | Integer  | false     | 0                                         | Adjust the date difference between the current date and the first date displayed in the GitHub contributions panel. E.g, toda is the 8th of March, but GitHubs panel is still one day behind(7th of March). So you enter -1 to change todays date to the one displayed on the panel.")                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| message         | String  | true    | hello                                             | Enter the message you would like to be displayed on the contribution graph. The maximum ammount of characters is 10. |

## How the letters work
Each letter/symbol/number is displayed in a 7x5 matrix. The GitHub contributions panel has space for 10 of those matrices.  
If you want to use a specific letter/symbol/number in your github message, you will have to check in the charMatrix.go folder and see if's already defined. Currently there are a few symbols and all the lower case letter. But don't worry. You can define other letters/numbers/symbols. Just define your symbol in a new matrix and add it to the alphabet variable. I've used this font to draw the letters: 
![pixel_fonts](https://user-images.githubusercontent.com/76880102/157720191-8197c37b-8f53-49c5-be57-78ab4ec8ddd2.jpg)


## How to run the random option
1. Create an fully empty GitHub repository (not even a README). Do not initialize it.
2. For the basic option (contribute the entire year), you run the command:
```sh
go run . random -repository="https://github.com/yourUserName/yourRepo.git"
```
Because it has to commit for each day of the year, this can take a few minutes to complete. <br />
Note: After the script finishes, you have to wait for GitHub to register and show the changes.

### Contributions graph before running the command above
![original_commit_graph](https://user-images.githubusercontent.com/76880102/155739686-1732adbe-4f4c-4639-b85c-47a1654e41e3.PNG)

### Contributions graph after committing on nearly each day of the yeah(depengind on the frequency parameter).
![graph after random script](https://user-images.githubusercontent.com/76880102/155740542-5f47f7cc-3eed-44aa-a031-fc51a16335ca.PNG)

Note: The default maximum value of commits in the example above was 7 per day. The script made a contribution 85% of days in a year.
### The month parameter
If you want your graph to look like you've made a lot of contributions in a specific month, you can do it like this:
```sh
go run . random -repository="https://github.com/yourUserName/yourRepo.git" -month=3-4 -frequency=96 -commit_limit=15
```

![graphh specific month 3-4](https://user-images.githubusercontent.com/76880102/155750671-3ec85410-9537-41ce-b100-7a49c437b0c3.PNG)
Lets say the current date is 25th of february. The command above will start making contributions from the 25 of march for 61 days (because the we specified -month=3-4 from march-april). And because we are nearly at the end of march the cript will continue past april, untill it commits the sum of days the entered months have.

You can also make different combinations. Just make another repository and run another command. E.g, you could run the default command whitch would post for the entire year, and then just make another empty repository and make more posts in a specific time, so it would look like you did something the entire year, but you were really productive at a specific time. 

## Paremeters
| Flag Name        | Type    | Required | Default                                        | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
|-------------------|---------|----------|------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| repository        | String  | true     |                                       | Enter a link to an empty non-initialized GitHub repository to which you want to push the generated file. The link can be an SSH (assuming you have an ssh key) or the HTTPS format. (e.g., git@github.com:yourusername/yourrepo.git or https://github.com/yourusername/yourrepo.git)                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| commit_limit          | Integer  | false     | 7                                         | Set the limit(n) of commits per single day. The script will randomly commit from 1 to n times a day. The maximum is 15 and the minimum is 1.                                                                                                                                                                                                                                                                                                      |
| frequency           | Integer  | false     | 75                                         | The procentage of days out of 365 you would like to contribute. E.g., if you enter 20, you will contribute 73 days out of 1 year.                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| month         | String  | false    | 1-12                                             | Contribute only in a specific period. If you enter 3-5, you will only commit from march(starting the same day of month as today) to may. Entering only one number like 8(october) will prompt the script to commit only on the specified month. <br /> <br /> Note that GitHubs contribution time frame is current date minus one year. So if our current month is february and you enter the flag -month=1-3, it won't contribute from january to march. It will contribute from february to april, because the 1st month in contributions panel is the current month. |

## Costumizations

You can do a lot of custom "stuff" here. You can run the random and the nonrandom option seperately to get more specific results. All you need to do it make a new repository on GitHub for each time you run the script. E.g, You run the random command with specific instructions and then just make another repository and run another version of the script there. 

If you dont like a specific result, you can just delete the repository and the commits will be removed.

## Having problems? 
Keep in mind that it can take GitHub several minutes to update the changes. Still no luck?

Ok sooooo, what could of went wront. Imma be honest here. I did not test this enough. But if you execute the script and it finishes, and there is still no change on the GitHub page, you will have to check your repository and see if there is a file named data.txt with a lot of commits. If you can't find it, you will have to make sure that you are logged in into the right GitHub account and that you have all the permissions to push to the repo. There is no error handling in this script (yet). I would recommend to just pushing a random file from the git bash. If it works, the script should work too.




