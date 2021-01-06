# Exam: DAT320 Operating Systems and Systems Programming

## Preamble

- Start: 01.12.20, 09:00
- End: 01.12.20, 13:00
- The exam will be posted on the course's `assignments` repository on GitHub.
- If the exam is found to contain errors or there is a need for clarification, such messages may be posted on the Discord `announcements` channel during the exam.

## Aids

All aids are permitted.
You are not allowed to get help from other people when working on your exam assignment.
We are also reminding you that you, when registering for the semester, signed that you have read and understood the rules for **cheating and plagiarism** in the Exam Rules and Regulations at the University of Stavanger.
Plagiarism control will be carried out.

## Important Contacts

If you need help during the exam, you can call the phone number below.
This applies if you need clarifications on the exam assignments.

- Course Responsible: Hein Meling, hein.meling@uis.no, Phone: 92436131

## Withdrawal During the Exam

If you wish to withdraw from the exam, you must follow these instructions:

- Create a file `withdraw.md` with the content `# deliver blank`
- Add the file to git and push it to GitHub.

These are the command line steps to withdraw:

```sh
cd exam
echo "# deliver blank" > withdraw.md
git add withdraw.md
git commit -m "withdraw"
git push
```

## Start of Exam

At the start of the exam, pull the `assignments` repository to retrieve the exam:

```sh
git pull course-assignments master
```

You can also view the exam at this URL [https://github.com/dat320-2020/assignments](https://github.com/dat320-2020/assignments).
The exam material can be found in the [`exam`](https://github.com/dat320-2020/assignments/tree/master/exam) folder.

## Exam Delivery

The exam **must** be delivered on your private repository within the `dat320-2020` organization on GitHub.
That is, the repository named `username-labs`, where `username` is your GitHub username.
The exam must be delivered (pushed to GitHub) at the end of the exam period.

- A 15 minute grace period will be given for the final delivery.
- It is recommended that you commit and push your changes several times during the exam.
- **VERY IMPORTANT**: Remember to **push** your final changes to GitHub.
  After the 15 minute grace period, no more changes will be accepted.

## Delivery Problems

If you have problems pushing your submission to GitHub, you must immediately send an email to hein.meling@uis.no with *information showing the last commit from the git log*.
See detailed instructions below.
This is an emergency solution, as a general rule all answers must be submitted to GitHub.

At the end of the exam period, make sure that your code has been committed locally:

```console
git status
git add <any changed files>
git commit
```

Then run the command:

```console
git log
```

Send the most recent log entry to hein.meling@uis.no, including the **commit hash**.
You are then expected to try to push again later, if the problem was a network connection problem.

```console
git push
```
