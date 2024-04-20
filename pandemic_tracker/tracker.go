package pandemic_tracker

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrUserNotRegistered     = errors.New("user not registered")
	ErrAdminNotRegistered    = errors.New("admin not registered")
)

const (
	ZoneCalcDuration = 15 * 24 * time.Hour // 15 days
	DefaultAdminId   = "admin1"
)

type User struct {
	Name       string
	Mobile     int
	PinCode    int
	isPositive bool
	recordedAt time.Time
}

type Tracker struct {
	Users          map[string]*User
	Admins         map[string]*Admin
	UsersByPinCode map[int]*PinCodeUserMap
}

type Admin struct {
	Name string
}

type PinCodeUserMap struct {
	users []*User
}

func NewTracker() *Tracker {
	Admins := make(map[string]*Admin)
	admin1 := &Admin{Name: DefaultAdminId}
	Admins[DefaultAdminId] = admin1
	return &Tracker{Users: make(map[string]*User), Admins: Admins, UsersByPinCode: make(map[int]*PinCodeUserMap)}
}

func (tracker *Tracker) RegisterUser(name string, mobile int, pinCode int) error {
	if tracker.Users[name] != nil {
		return ErrUserAlreadyRegistered
	}

	newUser := &User{
		Name:    name,
		Mobile:  mobile,
		PinCode: pinCode,
	}

	tracker.Users[name] = newUser

	if tracker.UsersByPinCode[pinCode] == nil {
		tracker.UsersByPinCode[pinCode] = &PinCodeUserMap{
			users: []*User{},
		}

	}

	tracker.UsersByPinCode[pinCode].users = append(tracker.UsersByPinCode[pinCode].users, newUser)

	fmt.Printf("user created successfully for %s\n", name)

	return nil
}

func (tracker *Tracker) SelfAssessment(userId string, symptoms []string, hasTravelHistory, hasContactPatient bool) (int, error) {
	user := tracker.Users[userId]
	if user == nil {
		return 0, ErrUserNotRegistered
	}

	if hasTravelHistory && hasContactPatient && len(symptoms) > 2 {
		return 95, nil
	} else if (len(symptoms) > 0 && hasTravelHistory) || hasContactPatient && hasTravelHistory || (len(symptoms) > 0 && hasContactPatient) {
		return 75, nil
	} else if len(symptoms) > 0 || hasTravelHistory || hasContactPatient {
		return 50, nil
	}

	return 5, nil
}

func (tracker *Tracker) PandemicResult(adminId string, userId string, isPositive bool, recordedAt time.Time) error {
	admin := tracker.Admins[adminId]
	if admin == nil {
		return ErrAdminNotRegistered
	}

	user := tracker.Users[userId]
	if user == nil {
		return ErrUserNotRegistered
	}

	user.isPositive = isPositive
	user.recordedAt = recordedAt

	if isPositive {
		fmt.Printf("Record of %s saved successfully as Positive", userId)
	} else {
		fmt.Printf("Record of %s saved successfully as Negative", userId)

	}
	return nil
}

func (tracker *Tracker) GetZone(adminId string, pinCode int) (string, error) {
	admin := tracker.Admins[adminId]
	if admin == nil {
		return "", ErrAdminNotRegistered
	}

	numPositiveCases := 0

	zone := tracker.UsersByPinCode[pinCode]

	if zone == nil {
		return "Green", nil
	}

	for _, user := range zone.users {
		if user.isPositive && user.recordedAt.Add(ZoneCalcDuration).After(time.Now()) {
			numPositiveCases++
		}
	}

	if numPositiveCases > 5 {
		return "Red", nil
	} else if numPositiveCases > 0 {
		return "Orange", nil
	}

	return "Green", nil
}
