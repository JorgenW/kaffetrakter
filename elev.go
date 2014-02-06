//#include "elev.h"
//#cgo CFLAGS: -std=c99
//#cgo LDFLAGS: -lcomedi - lm
import "C"


#define N_FLOORS 4

/**
  Button types for function elev_set_button_lamp() and elev_get_button().
*/
type tag_elev_lamp_type int
const ( 
    BUTTON_CALL_UP tag_elev_lamp_type = iota 
    BUTTON_CALL_DOWN
    BUTTON_COMMAND 
)



/**
  Initialize elevator.
  @return Non-zero on success, 0 on failure.
*/
func elev_init() int{
	return int(C.elev_init());
}


/**
  Sets the speed of the elevator. 
  @param speed New speed of elevator. Positive values denote upward movement
    and vice versa. Set speed to 0 to stop the elevator. From -300 to 300 gives
    sensible speeds. The hardware wil emit a constant tone if the speed is too high.
*/
func elev_set_speed(speed int) {
	C.elev_set_speed(int speed);
}


/**
  Get floor sensor signal.
  @return -1 if elevator is not on a floor. 0-3 if elevator is on floor. 0 is
    ground floor, 3 is top floor.
*/
func elev_get_floor_sensor_signal() int{
	return int(C.elev_get_floor_sensor_signal(void));
}


/**
  Gets a button signal.
  @param button Which button type to check. Can be BUTTON_CALL_UP,
    BUTTON_CALL_DOWN or BUTTON_COMMAND (button "inside the elevator).
  @param floor Which floor to check button. Must be 0-3.
  @return 0 if button is not pushed. 1 if button is pushed.
*/
func elev_get_button_signal(button elev_button_type_t, floor int) int{
	return int(C.elev_get_button_signal(elev_button_type_t button, int floor));
}


/**
  Get signal from stop button.
  @return 1 if stop button is pushed, 0 if not.
*/
func elev_get_stop_signal() int{
	return int(C.elev_get_stop_signal());
}


/**
  Get signal from obstruction switch.
  @return 1 if obstruction is enabled. 0 if not.
*/
func elev_get_obstruction_signal() int{
	return int(C.elev_get_obstruction_signal(void));
}


/**
  Set floor indicator lamp for a given floor.
  @param floor Which floor lamp to turn on. Other floor lamps are turned off.
*/
func elev_set_floor_indicator(floor int){
	C.elev_set_floor_indicator(int floor);
}


/**
  Set a button lamp.
  @param lamp Which type of lamp to set. Can be BUTTON_CALL_UP,
    BUTTON_CALL_DOWN or BUTTON_COMMAND (button "inside" the elevator).
  @param floor Floor of lamp to set. Must be 0-3
  @param value Non-zero value turns lamp on, 0 turns lamp off.
*/
func elev_set_button_lamp(button elev_button_type_t, floor int, value int) {
	C.elev_set_button_lamp(elev_button_type_t button, int floor, int value);
}


/**
  Turn stop lamp on or off.
  @param value Non-zero value turns lamp on, 0 turns lamp off.
*/
func elev_set_stop_lamp(value int){
	C.elev_set_stop_lamp(int value);
}


/**
  Turn door-open lamp on or off.
  @param value Non-zero value turns lamp on, 0 turns lamp off.
*/
func elev_set_door_open_lamp(value int){
	C.elev_set_door_open_lamp(int value);
}
