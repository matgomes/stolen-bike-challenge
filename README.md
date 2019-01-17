# Stolen Bike Cases - JOIN Coding Challenge - Backend (Node.js)
![JOIN Stolen Bike Cases](https://github.com/join-com/coding-challenge-backend-nodejs/raw/master/illustration.png)

## Context
Stolen bikes are a typical problem in Berlin. The Police want to be more efficient in resolving stolen bike cases. They decided to build a software that can automate their processes — the software that you're going to develop. 

## Product Requirements
- [ ] Bike owners can report a stolen bike.
- [ ] A bike can have multiple characteristics: license number, color, type, full name of the owner, date, and description of the theft.
- [ ] Police have multiple departments that are responsible for stolen bikes. 
- [ ] A department can have some amount of police officers who can work on stolen bike cases.
- [ ] The Police can scale their number of departments, and can increase the number of police officers per department.
- [ ] Each police officer should be able to search bikes by different characteristics in a database and see which department is responsible for a stolen bike case.
- [ ] New stolen bike cases should be automatically assigned to any free police officer in any department.  
- [ ] A police officer can only handle one stolen bike case at a time. 
- [ ] When the Police find a bike, the case is marked as resolved and the responsible police officer becomes available to take a new stolen bike case. 
- [ ] The system should be able to assign unassigned stolen bike cases automatically when a police officer becomes available.

## Your Mission
Your task is to provide APIs for a frontend application that satisfies all requirements above.

Please stick to the Product Requirements. You should not implement authorisation and authentication, as they are not important for the assessment. Assume everyone can make requests to any api. 

## Tech Requirements
- Node.js
- You are free to use any framework, but it’s recommended that you use one that you’re good at
- Use only SQL Database
- Tests (quality and coverage)
- Typescript is a plus

## Instructions
- Fork this repo
- The challenge is on!
- Build a performant, clean and well-structured solution
- Commit early and often. We want to be able to check your progress
- Make your API public. Deploy it using the service of your choice (e.g. AWS, Heroku, Digital Ocean...)
- Create a pull request
- Please complete your working solution within 7 days of receiving this challenge, and be sure to notify us when it is ready for review.
