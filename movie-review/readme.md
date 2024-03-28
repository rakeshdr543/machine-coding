Problem : Movie Review Platform

We have a requirement to implement a platform related to Movie reviews, and following is the description of it.

Platform Overview:
Movie review platform collects reviews for Movies from its users. Later these movie reviews are used to derive insights which helps in enriching the lives of its customers with entertainment.

Platform Capabilities:
● Adding Users and Movies.
○ Platform by default on-boards a user as a ‘viewer’.
○ A ‘viewer’ category can be upgraded to a ‘critic’ category after he/she published more than 3 reviews for various movies.
○ Critics are considered as experts in the judgement here, so critics reviews will be captured with more weightage. i.e. review rating of a critic will be considered twice.
● Users can review a movie
○ Users can only review Movies which have been released so far, they cannot review upcoming movies.
○ Users can give a review-score between 1 to 10. (Higher the score the better the liking for the movie). Currently we are not allowing a user to review the same Movie multiple times.
● User can Update/Delete their Review
● List all reviews given by a User
● List top n movies by total review score for a:
○ particular year of release
○ particular genre

Bonus Capabilities:
● Based on the users behaviour Platform should give capability to plugin more user upgradation policies from. Eg. Viewer ->Critic -> Expert
● List top n movies by total review score based on combination of pivots:
○ Pivots : [Year of release, User category, Genre]

Expectations:

Make sure that you have working and demonstrable code for all the above requirements.
Platform Capabilities should be strictly followed. Work on the expected output first and then move on the Bonus Capabilities.
Use of proper abstraction, separation of concerns is required.
Code should easily accommodate new requirements with minimal changes.
Proper exception handling is required.
Code should be modular, readable and unit-testable.
Important Notes:
● Do not use any database store, use in-memory data structure.
● Do not create any UI for the application.
● Do not build a Command line interface, Executing test cases or simple Main function should be sufficient
● Do not make any assumption, please ask it out.

Sample Test Cases:

Onboard 10 movies onto your platform in 3 different years.
a. Add Movie("Don" released in Year 2006 for Genres “Action” & “Comedy”)
b. Add Movie("Tiger" released in Year 2008 for Genre “Drama”)
c. Add Movie("Padmavat" released in Year 2006 for Genre “Comedy”)
d. Add Movie("Lunchbox-2" released in Year 2022 for Genre “Drama”)
e. Add Movie("Guru" released in Year 2006 for Genre “Drama”)
f. Add Movie("Metro" released in Year 2006 for Genre “Romance”)

Add users to the system:
a. Add User(“SRK”)
b. Add User(“Salman”)
c. Add User(“Deepika”)

Add/Update/Delete Reviews:
a. add_review(“SRK”, “Don”, 2)
b. add_review(“SRK”, “Padmavat”, 8)
c. add_review(“Salaman”, “Don”, 5)
d. add_review(“Deepika”, “Don”, 9)
e. add_review(“Deepika”, “Guru”, 6)
f. add_review(“SRK”,”Don”, 10) - Exception multiple reviews not allowed
g. update_review(“SRK”,”Don”, 8) - update won't affect the existing user category (viewer will remain viewer even after this operation)
h. add_review(“Deepika”, “Lunchbox-2”, 5) - Exception movie yet to be released
i. add_review(“SRK”, “Tiger”, 5). Since ‘SRK’ has published 3 reviews, he is promoted to ‘critic’ now.
j. delete_review(“SRK”, “Tiger”). Since ‘SRK’ has deleted his review he is downgraded back to viewer now, total active reviews given by SRK is < 3
k. add_review(“SRK”, “Metro”, 7) ‘SRK promoted back to critic’

List User Reviews:
a. list_review(“SRK”)
Output: {
“Don” : 8,
“Padmavat” : 8,
“Metro” : 7
}

List top 1 movie by review score in “2006” year:
a. Top in Year “2006”:
Output: Don - 9 + 5 + 8 \* 2 (at the time of evaluation SRK promoted to critic so rating considered twice)
Output: Don - 30 ratings (sum of all ratings of Deepika, Salman & SRK)
b. [Bonus eg] Top in Year “2006” by “critics_preferred”:
Output: Don - 8 ratings (SRK gave 7 for Metro as critic and 8 to Don )

List top 1 movie by review score in “Drama” genre:
Output: Guru - 6 ratings
