# GitHub contributions generator

A script written in Go, which goal is to make your GitHub contributions graph the way you want it to be.

## What's the goal? 
Let't transform this 
![original-github](https://user-images.githubusercontent.com/76880102/157703961-8ed845a2-4210-4b1c-b8fb-dbbeb3dd85ce.PNG)

to something more interesting

![bamboozled_med](https://user-images.githubusercontent.com/76880102/157708873-671f7c06-b5b2-490a-ad41-11173e450e2c.PNG)

## Disclaimer
This script was written for fun. In no way is it meant to "cheat". Nobody should judge a programmer regarding to his GitHub contributions graph.

### System requirements
This script is written in Go, so you will need it. And ofcourse you will need git. 

## How to run the nonrandom option (displaying a message)
1. Create an fully empty GitHub repository (not even a README). Do not initialize it.
2. For the basic option (contribute the entire year), you run the command:
```sh
go run . nonrandom -message="HelloWorld" -repository="https://github.com/yourUserName/yourRepo.git" -adjust_date=-3 -saturation=medium
```
Note the adjust_date parameter. It's a important one. More about it below. 

![helloworld_high](https://user-images.githubusercontent.com/76880102/157709681-304db8cb-e8b8-46b3-93a5-d37d4f40c698.PNG)

### The adjust_date parameter

If you want to display a message, you will have to check if the current date and the 1st date displayed in the GitHub contributions graph are the same. I have no idea how the graph works, but usually the first "box" (top left corner) is a few days behind the current date. 

![dateExxample](https://user-images.githubusercontent.com/76880102/157712749-0ab75ee0-01bf-4203-bc2d-979649414032.png)

In the example above we see that there is a 3 day difference between todays date and the first GitHub date. So we have to let the script know that It has to start 3 days earlier. So we use the -adjust_date=-3 flag. 

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
| commit_limit          | Integer  | false     | 7                                         | Set the limit(n) of commits per single day. The script will randomly commit from 1 to n times a day. The maximum is 15 and the minimum is 1.                                                                                                                                                             
                                                                                                                                    |
| frequency           | Integer  | false     | 75                                         | The procentage of days out of 365 you would like to contribute. E.g., if you enter 20, you will contribute 73 days out of 1 year.                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| month         | String  | false    | 1-12                                             | Contribute only in a specific period. If you enter 3-5, you will only commit from march(starting the same day of month as today) to may. Entering only one number like 8(october) will prompt the script to commit only on the specified month. <br /> <br /> Note that GitHubs contribution time frame is current date minus one year. So if our current month is february and you enter the flag -month=1-3, it won't contribute from january to march. It will contribute from february to april, because the 1st month in contributions panel is the current month. |



