import random
import string
from hstest import StageTest, dynamic_test, CheckResult, WrongAnswer, TestedProgram


class Test(StageTest):
    @dynamic_test
    def example_test(self):
        program = TestedProgram()

        reply = program.start()

        self.check_empty_or_none_output(reply)
        self.check_user_prompt(reply)

        reply = program.execute('create This is my first record!')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, 'create')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('create This is my second record!')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, 'create')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('list')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, 'list')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('exit')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Info] bye!')

        return CheckResult.correct()

    @dynamic_test
    def random_test(self):
        program = TestedProgram()

        program.start()

        command = self.get_random_string(5)
        command_argument = f'{self.get_random_string(6)} {self.get_random_string(10)}'

        reply = program.execute(f'{command} {command_argument}')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, command)

        return CheckResult.correct()

    @staticmethod
    def check_command_response(raw_output, message=''):
        output = raw_output.strip().lower().split("\n")[0]
        if not message and output:
            raise WrongAnswer(
                'The program should print nothing.\n'
                f'Your output: {raw_output}.'
            )
        if not output.strip().startswith(message.lower()):
            raise WrongAnswer(
                f'The program should print "{message}".\n'
                f'Your output: "{raw_output}".'
            )

    def check_user_prompt_after_command(self, raw_output):
        output = raw_output.strip().lower().split("\n")
        if len(output) < 2:
            raise WrongAnswer('The program should ask user for a command.')

        self.check_user_prompt(output[1])

    @staticmethod
    def check_user_prompt(raw_output):
        prompt = 'Enter a command and data:'
        if not raw_output.strip().lower().startswith(prompt.lower()) and raw_output != '':
            raise WrongAnswer(
                'The program should ask user for a command.\n'
                f'Your output should be equal to "{prompt}".\n'
                f'Your output: "{raw_output}".'
            )

    @staticmethod
    def check_empty_or_none_output(raw_output):
        if not raw_output:
            raise WrongAnswer('Your output is empty or None.')

    @staticmethod
    def get_random_string(length):
        letters = string.ascii_lowercase
        return ''.join(random.choice(letters) for _ in range(length))


if __name__ == '__main__':
    Test().run_tests()
