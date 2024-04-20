Pandemic Tracker
Build an application to calculate Pandemic risk ( e.g. Covid – we are 4 years too late for this ) for users where users can register and tell if they have any symptoms, travel history, or came in contact with any positive patient. Based on which application will calculate and return the risk percentage of the user.

1.  User:
    You need to provide following options for User:
    Register: A user can register by providing name, mobile number and pincode. Consider phone number as unique.

Self Assessment: Users can enter below data :
Any symptoms (cough, sore throat, etc). ( list )
Any Travel history in the last 15 days. ( in form of Yes and No )
Any contact with Pandemic positive patients. ( in form of Yes and No )

Assume symptoms be these :
cough
sore throat
headache
fever

Pandemic Risk: Users can see the risk of being affected by Pandemic based on self assessment.

Risk Calculation:
No symptoms, No travel history, No contact with Pandemic positive patient - Risk = 5%
Any one out of (symptoms, travel history or contact with Pandemic positive patient) is true - Risk = 50%
Any two out of (symptoms, travel history or contact with Pandemic positive patient) is true - Risk = 75%
More than 2 symptoms, travel history or contact with Pandemic positive patient is true - Risk = 95%

INPUT:

User capabilities
RegisterUser(ABC,9999999999,560037);

SelfAssessment(UserId, [ symptom 1 , symptom2 ] , Yes, No );
OUTPUT:

Risk is 75%

2. Admin:
   You need to provide Admin options for Pandemic Health workers:
   Register: Pandemic Health workers can register by mobile number and pin code on user’s behalf
   Pandemic Result: Health workers can enter the result of Pandemic tests for patients.
   Health workers can also mark already registered users as Recovered.

INPUT:

Admin Capabilities :

RegisterUser(ABC,9999999999,560037);
PandemicResult(AdminId, UserId1, Y);
PandemicResult(AdminId, UserId2, Y);
PandemicResult(AdminId, UserId2, N);

AdminId: Please take Admin’s name as AdminId. No requirement of having uniqueness check on admin name/id.

OUTPUT:

Record of UserId saved successfully as Positive

Or

Record of UserId saved successfully as Negative

3.  Zones:

Zone = pincode
Mark zones(pincodes) as green, orange and red based on positive Pandemic cases
Default zone - GREEN - <= 0 cases in that zone or no user has registered against that Zone.

> 0 and <=5 cases in a zone - ORANGE
> 5 cases in a zone - RED

Note : Please use only Admins inputs for calculating zones color

INPUT:
GetZone(560037);

OUTPUT:
No. of positive cases: 1
ORANGE

4. Bonus:
   Take DateTime as part of input in PandemicResult(AdminId, UserId1, Y) to be able to calculate the below result.

Change zones(pincodes) to green, orange and red based on positive Pandemic cases in last few days:
No cases in last 15 days in a zone - GREEN
<=5 cases in a zone in last 15 days - ORANGE

> 5 cases in a zone in last 15 days - RED

Expectations:
The code should be demo-able. Either by using the main driver program or test cases.
The code should handle edge cases properly and fail gracefully.
Create the sample data yourself. You can put it into a file, test case or main driver program itself. No extra points for creating UI or input parsing. Feel free to code it in the main method or test class.
Database integration is not required. Use in memory data structures instead.
A driver function with the main method should be sufficient. Please do not spend time building a spring boot or dropwizard application.
The programme should be functionally correct and complete.
The code should be modular. The code should have the basic OO design.
The code should be extensible. Wherever applicable, use interfaces and contracts between different methods. It should be easy to add/remove functionality without rewriting the entire codebase.
The code should be legible, readable and DRY.
Complete the mandatory tasks before moving to the bonus question.
Guidelines:
Please do not access the internet for anything EXCEPT syntax.
You are free to use the language and IDE of your choice.
The entire code should be your own.
