from hstest import StageTest, dynamic_test, CheckResult, WrongAnswer, TestedProgram


class Test(StageTest):
    @dynamic_test
    def test(self):
        program = TestedProgram()

        reply = program.start()

        self.check_empty_or_none_output(reply)
        self.check_number_of_notes_prompt(reply)

        reply = program.execute("3")
        self.check_empty_or_none_output(reply)
        self.check_user_prompt(reply)

        reply = program.execute('create')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Error] Missing note argument')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('del 2')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Error] Unknown command')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('list')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Info] Notepad is empty')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('create This is my first record!')
        self.check_empty_or_none_output(reply)
        self.check_command_response(
            reply, '[OK] The note was successfully created')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('list')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Info] 1: This is my first record!')
        self.check_user_prompt_after_command(reply)

        notes = [
            'This is my second record!',
            'This is my third record!'
        ]

        for note in notes:
            reply = program.execute(f'create {note}')
            self.check_empty_or_none_output(reply)
            self.check_command_response(
                reply, '[OK] The note was successfully created')
            self.check_user_prompt_after_command(reply)

        reply = program.execute(f'create This is my sixth record!')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Error] Notepad is full')
        self.check_user_prompt_after_command(reply)

        reply = program.execute(f'clear')
        self.check_empty_or_none_output(reply)
        self.check_command_response(
            reply, '[OK] All notes were successfully deleted')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('list')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Info] Notepad is empty')
        self.check_user_prompt_after_command(reply)

        reply = program.execute('exit')
        self.check_empty_or_none_output(reply)
        self.check_command_response(reply, '[Info] bye!')

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
    def check_number_of_notes_prompt(raw_output):
        prompt = 'Enter the maximum number of notes:'
        if not raw_output.strip().lower().startswith(prompt.lower()):
            raise WrongAnswer(
                'The program should ask user to enter the maximum number of notes\n'
                f'Your output should be equal to "{prompt}".\n'
                f'Your output: "{raw_output}".'
            )

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


if __name__ == '__main__':
    Test().run_tests()
