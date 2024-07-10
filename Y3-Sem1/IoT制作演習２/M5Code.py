from m5stack import *
from m5ui import *
from uiflow import *
import time
import random
import network
import urequests

remoteInit()
env3_0 = unit.get(unit.ENV3, unit.PORTA)
pir0 = unit.get(unit.PIR, unit.PORTB)

x = None
hours = None
secs = None
mins = None
is_paused = True
in_main_screen = False
last_time = None
last_motion_time = None  # Variable to track the last motion detection time

# APIから時刻を取得する関数
def get_japan_time():
    response = urequests.get('https://worldtimeapi.org/api/timezone/Asia/Tokyo')
    if response.status_code == 200:
        data = response.json()
        datetime_str = data['datetime']
        year = int(datetime_str[0:4])
        month = int(datetime_str[5:7])
        day = int(datetime_str[8:10])
        hour = int(datetime_str[11:13])
        minute = int(datetime_str[14:16])
        second = int(datetime_str[17:19])
        return (year, month, day, hour, minute, second)
    else:
        return None

def display_current_japan_time():
    japan_time = get_japan_time()
    if japan_time:
        year, month, day, hour, minute, second = japan_time
        time_str = "{:02}:{:02}".format(hour, minute)
        date_str = "{:04}-{:02}-{:02}".format(year, month, day)
        return (time_str, date_str)

lcd.clear()
lcd.font(lcd.FONT_DejaVu24)
lcd.setCursor(20, 20)

def show_main_screen():
    global in_main_screen
    in_main_screen = True
    setScreenColor(0xEA9999)
    nowtime = display_current_japan_time()
    if nowtime:
        t, d = nowtime
        nt = "{} | {}".format(d, t)
        T_back = M5TextBox(20, 10, nt, lcd.FONT_DejaVu18, 0xFFFFFF, rotate=0)
    lcd.qrcode('http://flow-remote.m5stack.com/?remote=undefined', 85, 40, 150)
    m_timer = M5TextBox(40, 210, "Timer", lcd.FONT_DejaVu18, 0xFFFFFF, rotate=0)

def blackout_screen():
    lcd.clear()

def _remote_ON_OFF(x):
    if x == 1:
        rgb.setColorAll(0x3333ff)
    else:
        rgb.setColorAll(0x000000)

def _remote_Bright(x):
    rgb.setBrightness(x)

def buttonA_wasPressed():
    global hours, secs, mins, is_paused, in_main_screen, last_time
    if in_main_screen:
        show_timer_screen()
    else:
        # Reset the timer
        hours, mins, secs = 0, 0, 0
        is_paused = True
        Start.setText("Start")
        last_time = None
        # Turn off LED
        rgb.setColorAll(0x000000)
        update_display()
btnA.wasPressed(buttonA_wasPressed)

def buttonB_wasPressed():
    global is_paused, last_time
    if not in_main_screen:
        is_paused = not is_paused
        Start.setText("Stop" if not is_paused else "Start")
        if is_paused:
            # Turn off LED
            rgb.setColorAll(0x000000)
        else:
            # Initial LED state for start
            rgb.setColorAll(0xFFFFFF)
            last_time = time.time()
btnB.wasPressed(buttonB_wasPressed)

def buttonC_wasPressed():
    global is_paused, hours, mins, secs, last_time
    if not in_main_screen:
        is_paused = True
        hours, mins, secs = 0, 0, 0
        is_paused = True
        Start.setText("Start")
        last_time = None
        rgb.setColorAll(0x000000)
        show_main_screen()
btnC.wasPressed(buttonC_wasPressed)

# Timer
def show_timer_screen():
    global in_main_screen, TitleBG, Timer, Hours, Hour, Min, label3, Secs, Start, Reset, T_back, bar_top, bar_bottom
    in_main_screen = False
    setScreenColor(0x222222)
    TitleBG = M5Rect(0, 0, 320, 40, 0x7f7fff)
    Timer = M5TextBox(125, 10, "TIMER", lcd.FONT_DejaVu24, 0x000000, rotate=0)
    Hour = M5TextBox(40, 116, "00", lcd.FONT_DejaVu40, 0xFFFFFF, rotate=0)
    Min = M5TextBox(130, 116, "00", lcd.FONT_DejaVu40, 0xFFFFFF, rotate=0)
    Secs = M5TextBox(222, 116, "00", lcd.FONT_DejaVu40, 0xFFFFFF, rotate=0)
    Hours = M5TextBox(108, 114, ":", lcd.FONT_DejaVu40, 0xFFFFFF, rotate=0)
    label3 = M5TextBox(194, 114, ":", lcd.FONT_DejaVu40, 0xFFFFFF, rotate=0)
    Start = M5TextBox(135, 210, "Start", lcd.FONT_DejaVu18, 0xFFFFFF, rotate=0)
    Reset = M5TextBox(40, 210, "Reset", lcd.FONT_DejaVu18, 0xFFFFFF, rotate=0)
    T_back = M5TextBox(230, 210, "back", lcd.FONT_DejaVu18, 0xFFFFFF, rotate=0)
    bar_top = M5Rect(0, 100, 320, 4, 0xFFFFFF)
    bar_bottom = M5Rect(0, 160, 320, 4, 0xFFFFFF)
    update_display()

def update_display():
    global Hour, Min, Secs, hours, mins, secs
    hours_str = str(hours) if hours >= 10 else '0' + str(hours)
    mins_str = str(mins) if mins >= 10 else '0' + str(mins)
    secs_str = str(secs) if secs >= 10 else '0' + str(secs)

    Hour.setText(hours_str)
    Min.setText(mins_str)
    Secs.setText(secs_str)

hours = 0
mins = 0
secs = 0

show_main_screen()  # Initialize screen

while True:
    current_time = time.time()
    
    # Check PIR sensor
    if pir0.state == 1:
        last_motion_time = current_time
        show_main_screen()
    elif last_motion_time and current_time - last_motion_time >= 3:
        blackout_screen()
        last_motion_time = None  # Reset to ensure screen stays black until motion is detected again

    if not is_paused:
        if last_time is not None:
            elapsed = current_time - last_time
            if elapsed >= 1:
                secs += int(elapsed)
                last_time = current_time
                if secs >= 60:
                    mins += secs // 60
                    secs = secs % 60
                    if mins >= 60:
                        hours += mins // 60
                        mins = mins % 60
                        if hours >= 24:
                            hours = 0
                            mins = 0
                            secs = 0

    if not in_main_screen:
        update_display()

    if secs % 30 == 0 and secs != 0:
        speaker.tone(1000, 200) 

    # LED flashing effect
    if not is_paused:
        R = random.randint(0, 255)
        G = random.randint(0, 255)
        B = random.randint(0, 255)
        for i in range(0, 256, 10):
            rgb.setColorFrom(6, 10, (R << 16) | (G << 8) | B)
            rgb.setColorFrom(1, 5, (G << 16) | (R << 8) | B)
            rgb.setBrightness(i)
            wait_ms(2)
        for i in range(255, -1, -10):
            rgb.setColorFrom(6, 10, (R << 16) | (G << 8) | B)
            rgb.setColorFrom(1, 5, (G << 16) | (R << 8) | B)
            rgb.setBrightness(i)
            wait_ms(2)

    wait_ms(2)
