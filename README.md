# GitHub contributions generator

A script written in Go, which goal is to make your GitHub contributions graph the way you want it to be.

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


