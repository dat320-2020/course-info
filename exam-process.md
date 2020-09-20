# Exam Information and Procedures

These are tentative instructions and may be revised before the exam.
If you have feedback or suggested improvements, please reach out to me.

The exam in DAT320 will be conducted as home exam using GitHub.
The duration of the exam will be 4 hours.

The exam will be posted on the `assignments` repository.
The exam must be delivered on your private repository.
That is, the repository named `username-labs`, where `username` is your GitHub username.
This is to reuse your existing repository and avoid potential issues with creating a new repository.

The exam must be delivered (pushed to GitHub) at the end of the exam period.
We will allow a short grace period to account for technical issues.
See also [Backup Plan](#backup-plan) below.

## Exam Content

The exam tasks will mainly be Go programming, following a similar approach to that of the lab assignments.
My objective will be to develop small simulated pieces of an operating system, where I will draw inspiration from the text book syllabus.
I may also use the existing lab assignments as template for you to build on, so it may payoff to work extra hard to understand the lab assignments.

During the exam, you will not have the benefit of Autograder's feedback, nor will we provide local tests.
We will develop and use Go tests to be used to check your code, and calculate your grade.
To that end, it is important that you follow the instructions given in the exam and in any code templates that we provide.
For example, if you change a function signature to something unexpected, it may break our checks, which could negatively impact your grade.
Obviously, your code must compile and it must be formatted according to the `gofmt` tool.
We may provide some tests to help ensure appropriate formatting.
If you submit code that does not compile, you will receive the grade F.

Here are some tips for preparing for the exam.

1. Ensure that your laptop is correctly configured to work with GitHub and your preferred code editor.
   If you have access to two separate machines, e.g. a stationary, consider setting up both with an equal work environment, as a backup.

2. Ensure that you are familiar with the existing lab assignments code, as you may be expected to extend it.

3. The exam period is short, so it will payoff to spend time practice coding and working with git and GitHub.

## Backup Plan

If you experience technical difficulties accessing GitHub, here is what you do.

### At Start of Exam

If you are unable to access the exam via GitHub send an email to hein.meling@uis.no.
I will then provide a copy of the `assignments` repository.

### At End of Exam Period

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

Send the most recent log entry to hein.meling@uis.no.
You are then expected to try to push again later.

```console
git push
```
