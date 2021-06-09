# codeing=UTF-8
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.common.by import By
import threading
import time


def join_game():
    browser = webdriver.Chrome()
    browser.implicitly_wait(8)
    browser.get('http://localhost:3000/game/test1')
    ready_btn = browser.find_element_by_xpath('//*[@id="root"]/div/div/div/button[2]')
    ready_btn.click()

    browser.implicitly_wait(8)
    calllord_btn = browser.find_element_by_xpath('//*[@id="root"]/div/div[4]/div/button[1]')
    calllord_btn.click()


    time.sleep(90)


if __name__ == "__main__":
    threadNum = 3
    threads = []
    cur = 0
    while cur < threadNum:
        thread = threading.Thread(target=join_game,name='thread_' + str(cur))
        threads.append(thread)
        cur+=1

    for thread in threads:
        thread.start()

    for thread in threads:
        thread.join()

    

