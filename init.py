import datetime
import os
import requests

# As of now only Java is supoorted 
# but hey, contributions are always welcomed

TEMPLATE_JAVA_CODE_URL = 'https://raw.githubusercontent.com/OBrutus/prepare-code/mainline/template.java'

if 'template.java' not in os.listdir():
    print('Taking the code from cloud\n')
    TEMPLATE_JAVA_CODE = requests.get(TEMPLATE_JAVA_CODE_URL).text
else:
    print('Using local template\n')
    template_reader = open('template.java', 'r')
    TEMPLATE_JAVA_CODE = template_reader.read()
    template_reader.close()

EXTENSION = '.java'

CODEFORCES = 'codeforces'

PLATFORM_SET = set(['codeforces', 'codechef', 'leetcode', 'interviewbit', 'hackerearth', 'hackerrank'])
def get_formatted_date():
    current_date = datetime.datetime.now()
    return current_date.strftime("%y.%m.%d")

FORMATTED_DATE = get_formatted_date()

print('''
    Hi, nice choice!!
    A very good day to code
    Today's date :: ''' + FORMATTED_DATE + '''
    ''')

# setting problem link
PROBLEM_LINK = input('''
        Please enter the problem link :: ''')
if PROBLEM_LINK[-1] == '/':
    PROBLEM_LINK = PROBLEM_LINK[:-1]

# setting platform
PLATFORM = 'Misc.'
for platform in PLATFORM_SET:
    if PROBLEM_LINK.find(platform) != -1:
        print('''
    Platform chose is :: ''' + platform + '''
                ''')
        PLATFORM = platform

PLACEHOLDER_MAPPINGS = {
    '{USERNAME}': os.getlogin(),
    '{DATE}': FORMATTED_DATE,
    '{PROBLEM_LINK}': PROBLEM_LINK,
    '{PLATFORM}': PLATFORM
}

# swap placeholders
for placeholder in PLACEHOLDER_MAPPINGS.keys():
    TEMPLATE_JAVA_CODE = TEMPLATE_JAVA_CODE.replace(placeholder, PLACEHOLDER_MAPPINGS[placeholder])

CONTEST_NUMBER = ''
PROBLEM_NUMBER = ''
# getting the file name
def get_file_name():
    if PLATFORM == CODEFORCES:
        # sample format :: https://codeforces.com/contest/{CONTEST_NUMBER}/problem/{PROBLEM_NUMBER}
        contest_format = 'contest/'
        index = PROBLEM_LINK.index(contest_format)
        index += len(contest_format)
        end_index = index + PROBLEM_LINK[index:].index('/')
        CONTEST_NUMBER = PROBLEM_LINK[index : end_index]
        PROBLEM_NUMBER = PROBLEM_LINK[ PROBLEM_LINK.rindex('/') + 1 :  ]
        return CONTEST_NUMBER + '-' + PROBLEM_NUMBER + EXTENSION
    
    FILE_NAME = input('''
            Enter the file name :: ''')
    return FILE_NAME + EXTENSION

FILE_NAME = get_file_name()
DESTINATION_DIR = PLATFORM + '/' + CONTEST_NUMBER + '/'

if not os.path.isdir(DESTINATION_DIR):
    print('DEBUG: creating ' + DESTINATION_DIR)
    os.makedirs(DESTINATION_DIR)

working_file = open(DESTINATION_DIR + FILE_NAME, 'w')
working_file.write(TEMPLATE_JAVA_CODE)
working_file.close()
print('''
        You can go to the path :: ''' + DESTINATION_DIR + '''
        File exist by name :: ''' + FILE_NAME + '''
        
        Opening via VSCode
        ''')

# Opening VSCode here, if not a fan of VSCode change this
os.system('code ' + DESTINATION_DIR)
