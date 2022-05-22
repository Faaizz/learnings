# Software Architecture Foundations

## Content
- [Software Architecture Foundations](#software-architecture-foundations)
  - [Content](#content)
  - [Design Process](#design-process)
    - [Define the problem](#define-the-problem)
    - [Develop User Stories](#develop-user-stories)
    - [Define the program structure](#define-the-program-structure)

## Design Process
Basic steps:
1. Define the problem
2. Develop User Stories
3. Define the structure

### Define the problem
In defining the problem, we want to develop a problem statement. 
Our problem statement should include the "problem" and a "solution".
Both problem and solution must be specified from a domain perspective.

### Develop User Stories
A user story should contain a domain actor, a domain action, and a domain value.
If a user story is too big (complicated), we can refine it by narrowing.
An example of narrowing is workflow isolation. 
In workflow isolation, we layout all the different workflow contained in the user story (e.g. with a UML activity diagram) and pick any single path form start to finish (while ignoring the other paths).

### Define the program structure
Here, we pick up a user story and perform the following actions:
1. Identify the events
2. Pinpoint the actions associated with the identified events
3. Find out the agent (and entity within the agent) that perform those actions
4. Relate the agent's action with a bounded context

A simplified example:
- Event: comment posted
- Action: post comment
- Agent: reader
- Bounded Context: commenting
