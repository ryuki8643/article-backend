drop table if exists codes;
drop table if exists steps;
drop table if exists articles;
create table if not exists articles
(
    article_id int,
    title text default 'no_title',
    author text default 'no_author',
    likes int default 0,
    primary key (article_id)
    );

create table if not exists steps
(
    step_primary_key int,
    article_id int,
    step_id int,
    step_title text,
    article_content text,
    foreign key (article_id) references articles(article_id),
    primary key (step_primary_key)
    );

create table if not exists codes
(
    step_primary_key int,
    code_id int,
    code_file_name text default 'no_file',
    code_content text default 'no_content',

    foreign key (step_primary_key) references steps(step_primary_key),
    primary key (step_primary_key,code_id)
    );
insert into articles values (0,'A first splash into JavaScript','mdn web docs',0);
insert into steps values (0,0,0,'Initial setup','To begin this tutorial, we''d like you to make a local copy of the number-guessing-game-start.html file (see it live here). Open it in both your text editor and your web browser. At the moment you''ll see a simple heading, paragraph of instructions and form for entering a guess, but the form won''t currently do anything.

The place where we''ll be adding all our code is inside the <script> element at the bottom of the HTML:' ||
                                                '    this content is from mdn web docs
https://developer.mozilla.org/en-US/docs/Learn/JavaScript/First_steps/A_first_splash
'),
                      (1,0,1,'Adding variables to store our data','Let''s get started. First of all, add the following lines inside your <script> element:This section of the code sets up the variables and constants we need to store the data our program will use.

Variables are basically names for values (such as numbers, or strings of text). You create a variable with the keyword let followed by a name for your variable.

Constants are also used to name values, but unlike variables, you can''t change the value once set. In this case, we are using constants to store references to parts of our user interface. The text inside some of these elements might change, but each constant always references the same HTML element that it was initialized with. You create a constant with the keyword const followed by a name for the constant.

You can assign a value to your variable or constant with an equals sign (=) followed by the value you want to give it.    this content is from mdn web docs
https://developer.mozilla.org/en-US/docs/Learn/JavaScript/First_steps/A_first_splash
'),
                      (2,0,2,'Functions','Next, add the following below your previous JavaScript:' ||
                                             'Functions are reusable blocks of code that you can write once and run again and again, saving the need to keep repeating code all the time. This is really useful. There are a number of ways to define functions, but for now we''ll concentrate on one simple type. Here we have defined a function by using the keyword function, followed by a name, with parentheses put after it. After that, we put two curly braces ({ }). Inside the curly braces goes all the code that we want to run whenever we call the function.

When we want to run the code, we type the name of the function followed by the parentheses.

Let''s try that now. Save your code and refresh the page in your browser. Then go into the developer tools JavaScript console, and enter the following line:
checkGuess();
After pressing Return/Enter, you should see an alert come up that says I am a placeholder; we have defined a function in our code that creates an alert whenever we call it.  this content is from mdn web docs
https://developer.mozilla.org/en-US/docs/Learn/JavaScript/First_steps/A_first_splash
'),
                      (3,0,3,'Conditionals','Returning to our checkGuess() function, I think it''s safe to say that we don''t want it to just spit out a placeholder message. We want it to check whether a player''s guess is correct or not, and respond appropriately.

At this point, replace your current checkGuess() function with this version instead:' ||
                                            'This is a lot of code — phew! Let''s go through each section and explain what it does.

The first line declares a variable called userGuess and sets its value to the current value entered inside the text field. We also run this value through the built-in Number() constructor, just to make sure the value is definitely a number. Since we''re not changing this variable, we''ll declare it using const.
Next, we encounter our first conditional code block. A conditional code block allows you to run code selectively, depending on whether a certain condition is true or not. It looks a bit like a function, but it isn''t. The simplest form of conditional block starts with the keyword if, then some parentheses, then some curly braces. Inside the parentheses, we include a test. If the test returns true, we run the code inside the curly braces. If not, we don''t, and move on to the next bit of code. In this case, the test is testing whether the guessCount variable is equal to 1 (i.e. whether this is the player''s first go or not):
guessCount === 1
Copy to Clipboard
If it is, we make the guesses paragraph''s text content equal to Previous guesses:. If not, we don''t.
Line 6 appends the current userGuess value onto the end of the guesses paragraph, plus a blank space so there will be a space between each guess shown.
The next block does a few checks:
The first if (){ } checks whether the user''s guess is equal to the randomNumber set at the top of our JavaScript. If it is, the player has guessed correctly and the game is won, so we show the player a congratulations message with a nice green color, clear the contents of the Low/High guess information box, and run a function called setGameOver(), which we''ll discuss later.
Now we''ve chained another test onto the end of the last one using an else if (){ } structure. This one checks whether this turn is the user''s last turn. If it is, the program does the same thing as in the previous block, except with a game over message instead of a congratulations message.
The final block chained onto the end of this code (the else { }) contains code that is only run if neither of the other two tests returns true (i.e. the player didn''t guess right, but they have more guesses left). In this case we tell them they are wrong, then we perform another conditional test to check whether the guess was higher or lower than the answer, displaying a further message as appropriate to tell them higher or lower.
The last three lines in the function (lines 26–28 above) get us ready for the next guess to be submitted. We add 1 to the guessCount variable so the player uses up their turn (++ is an incrementation operation — increment by 1), and empty the value out of the form text field and focus it again, ready for the next guess to be entered.    this content is from mdn web docs
https://developer.mozilla.org/en-US/docs/Learn/JavaScript/First_steps/A_first_splash
'),
                      (4,0,4,'Events','At this point, we have a nicely implemented checkGuess() function, but it won''t do anything because we haven''t called it yet. Ideally, we want to call it when the "Submit guess" button is pressed, and to do this we need to use an event. Events are things that happen in the browser — a button being clicked, a page loading, a video playing, etc. — in response to which we can run blocks of code. Event listeners observe specific events and call event handlers, which are blocks of code that run in response to an event firing.

Add the following line below your checkGuess() function:

guessSubmit.addEventListener(''click'', checkGuess);
Copy to Clipboard
Here we are adding an event listener to the guessSubmit button. This is a method that takes two input values (called arguments) — the type of event we are listening out for (in this case click) as a string, and the code we want to run when the event occurs (in this case the checkGuess() function). Note that we don''t need to specify the parentheses when writing it inside addEventListener().

Try saving and refreshing your code now, and your example should work — to a point. The only problem now is that if you guess the correct answer or run out of guesses, the game will break because we''ve not yet defined the setGameOver() function that is supposed to be run once the game is over. Let''s add our missing code now and complete the example functionality.  this content is from mdn web docs
https://developer.mozilla.org/en-US/docs/Learn/JavaScript/First_steps/A_first_splash
'),
                      (5,0,5,'Finishing the game functionality','Let''s add that setGameOver() function to the bottom of our code and then walk through it. Add this now, below the rest of your JavaScript:The first two lines disable the form text input and button by setting their disabled properties to true. This is necessary, because if we didn''t, the user could submit more guesses after the game is over, which would mess things up.
The next three lines generate a new <button> element, set its text label to "Start new game", and add it to the bottom of our existing HTML.
The final line sets an event listener on our new button so that when it is clicked, a function called resetGame() is run.
Now we need to define this function too! Add the following code, again to the bottom of your JavaScript:This rather long block of code completely resets everything to how it was at the start of the game, so the player can have another go. It:

Puts the guessCount back down to 1.
Empties all the text out of the information paragraphs. We select all paragraphs inside <div class="resultParas"></div>, then loop through each one, setting their textContent to '''' (an empty string).
Removes the reset button from our code.
Enables the form elements, and empties and focuses the text field, ready for a new guess to be entered.
Removes the background color from the lastResult paragraph.
Generates a new random number so that you are not just guessing the same number again!
At this point, you should have a fully working (simple) game — congratulations!

All we have left to do now in this article is to talk about a few other important code features that you''ve already seen, although you may have not realized it.    this content is from mdn web docs
https://developer.mozilla.org/en-US/docs/Learn/JavaScript/First_steps/A_first_splash
');
insert into codes values (0,0,'index.html','<!DOCTYPE html>
<html lang="en-us">
<head>
    <meta charset="utf-8">

    <title>Number guessing game</title>

    <style>
        html {
            font-family: sans-serif;
        }

        body {
            width: 50%;
            max-width: 800px;
            min-width: 480px;
            margin: 0 auto;
        }

        .form input[type="number"] {
            width: 200px;
        }

        .lastResult {
            color: white;
            padding: 3px;
        }
    </style>
</head>

<body>
    <h1>Number guessing game</h1>

    <p>We have selected a random number between 1 and 100. See if you can guess it in 10 turns or fewer. We''ll tell you if your guess was too high or too low.</p>

    <div class="form">
        <label for="guessField">Enter a guess: </label>
        <input type="number" min="1" max="100" required id="guessField" class="guessField">
        <input type="submit" value="Submit guess" class="guessSubmit">
    </div>

    <div class="resultParas">
        <p class="guesses"></p>
        <p class="lastResult"></p>
        <p class="lowOrHi"></p>
    </div>

    <script>

        // Your JavaScript goes here

    </script>
</body>
</html>'),(1,0,'index.html','<!DOCTYPE html>
<html lang="en-us">
<head>
    <meta charset="utf-8">

    <title>Number guessing game</title>

    <style>
        html {
            font-family: sans-serif;
        }

        body {
            width: 50%;
            max-width: 800px;
            min-width: 480px;
            margin: 0 auto;
        }

        .form input[type="number"] {
            width: 200px;
        }

        .lastResult {
            color: white;
            padding: 3px;
        }
    </style>
</head>

<body>
    <h1>Number guessing game</h1>

    <p>We have selected a random number between 1 and 100. See if you can guess it in 10 turns or fewer. We''ll tell you if your guess was too high or too low.</p>

    <div class="form">
        <label for="guessField">Enter a guess: </label>
        <input type="number" min="1" max="100" required id="guessField" class="guessField">
        <input type="submit" value="Submit guess" class="guessSubmit">
    </div>

    <div class="resultParas">
        <p class="guesses"></p>
        <p class="lastResult"></p>
        <p class="lowOrHi"></p>
    </div>

    <script>
        let randomNumber = Math.floor(Math.random() * 100) + 1;

        const guesses = document.querySelector(''.guesses'');
        const lastResult = document.querySelector(''.lastResult'');
        const lowOrHi = document.querySelector(''.lowOrHi'');

        const guessSubmit = document.querySelector(''.guessSubmit'');
        const guessField = document.querySelector(''.guessField'');

        let guessCount = 1;
        let resetButton;
    </script>
</body>
</html>
'),(2,0,'index.html','<!DOCTYPE html>
<html lang="en-us">
<head>
    <meta charset="utf-8">

    <title>Number guessing game</title>

    <style>
        html {
            font-family: sans-serif;
        }

        body {
            width: 50%;
            max-width: 800px;
            min-width: 480px;
            margin: 0 auto;
        }

        .form input[type="number"] {
            width: 200px;
        }

        .lastResult {
            color: white;
            padding: 3px;
        }
    </style>
</head>

<body>
    <h1>Number guessing game</h1>

    <p>We have selected a random number between 1 and 100. See if you can guess it in 10 turns or fewer. We''ll tell you if your guess was too high or too low.</p>

    <div class="form">
        <label for="guessField">Enter a guess: </label>
        <input type="number" min="1" max="100" required id="guessField" class="guessField">
        <input type="submit" value="Submit guess" class="guessSubmit">
    </div>

    <div class="resultParas">
        <p class="guesses"></p>
        <p class="lastResult"></p>
        <p class="lowOrHi"></p>
    </div>

    <script>
        let randomNumber = Math.floor(Math.random() * 100) + 1;

        const guesses = document.querySelector(''.guesses'');
        const lastResult = document.querySelector(''.lastResult'');
        const lowOrHi = document.querySelector(''.lowOrHi'');

        const guessSubmit = document.querySelector(''.guessSubmit'');
        const guessField = document.querySelector(''.guessField'');

        let guessCount = 1;
        let resetButton;
        function checkGuess() {
            alert(''I am a placeholder'');
        }
        checkGuess();
    </script>
</body>
</html>
'),(3,0,'index.html','<!DOCTYPE html>
<html lang="en-us">
<head>
    <meta charset="utf-8">

    <title>Number guessing game</title>

    <style>
        html {
            font-family: sans-serif;
        }

        body {
            width: 50%;
            max-width: 800px;
            min-width: 480px;
            margin: 0 auto;
        }

        .form input[type="number"] {
            width: 200px;
        }

        .lastResult {
            color: white;
            padding: 3px;
        }
    </style>
</head>

<body>
    <h1>Number guessing game</h1>

    <p>We have selected a random number between 1 and 100. See if you can guess it in 10 turns or fewer. We''ll tell you if your guess was too high or too low.</p>

    <div class="form">
        <label for="guessField">Enter a guess: </label>
        <input type="number" min="1" max="100" required id="guessField" class="guessField">
        <input type="submit" value="Submit guess" class="guessSubmit">
    </div>

    <div class="resultParas">
        <p class="guesses"></p>
        <p class="lastResult"></p>
        <p class="lowOrHi"></p>
    </div>

    <script>
        let randomNumber = Math.floor(Math.random() * 100) + 1;

        const guesses = document.querySelector(''.guesses'');
        const lastResult = document.querySelector(''.lastResult'');
        const lowOrHi = document.querySelector(''.lowOrHi'');

        const guessSubmit = document.querySelector(''.guessSubmit'');
        const guessField = document.querySelector(''.guessField'');

        let guessCount = 1;
        let resetButton;
        function checkGuess() {
            const userGuess = Number(guessField.value);
            if (guessCount === 1) {
                guesses.textContent = ''Previous guesses: '';
            }
            guesses.textContent += `${userGuess} `;

            if (userGuess === randomNumber) {
                lastResult.textContent = ''Congratulations! You got it right!'';
                lastResult.style.backgroundColor = ''green'';
                lowOrHi.textContent = '''';
                setGameOver();
            } else if (guessCount === 10) {
                lastResult.textContent = ''!!!GAME OVER!!!'';
                lowOrHi.textContent = '''';
                setGameOver();
            } else {
                lastResult.textContent = ''Wrong!'';
                lastResult.style.backgroundColor = ''red'';
                if (userGuess < randomNumber) {
                    lowOrHi.textContent = ''Last guess was too low!'';
                } else if (userGuess > randomNumber) {
                    lowOrHi.textContent = ''Last guess was too high!'';
                }
            }

            guessCount++;
            guessField.value = '''';
            guessField.focus();
        }
        checkGuess();
    </script>
</body>
</html>
'),(4,0,'index.html','<!DOCTYPE html>
<html lang="en-us">
<head>
    <meta charset="utf-8">

    <title>Number guessing game</title>

    <style>
        html {
            font-family: sans-serif;
        }

        body {
            width: 50%;
            max-width: 800px;
            min-width: 480px;
            margin: 0 auto;
        }

        .form input[type="number"] {
            width: 200px;
        }

        .lastResult {
            color: white;
            padding: 3px;
        }
    </style>
</head>

<body>
    <h1>Number guessing game</h1>

    <p>We have selected a random number between 1 and 100. See if you can guess it in 10 turns or fewer. We''ll tell you if your guess was too high or too low.</p>

    <div class="form">
        <label for="guessField">Enter a guess: </label>
        <input type="number" min="1" max="100" required id="guessField" class="guessField">
        <input type="submit" value="Submit guess" class="guessSubmit">
    </div>

    <div class="resultParas">
        <p class="guesses"></p>
        <p class="lastResult"></p>
        <p class="lowOrHi"></p>
    </div>

    <script>
        let randomNumber = Math.floor(Math.random() * 100) + 1;

        const guesses = document.querySelector(''.guesses'');
        const lastResult = document.querySelector(''.lastResult'');
        const lowOrHi = document.querySelector(''.lowOrHi'');

        const guessSubmit = document.querySelector(''.guessSubmit'');
        const guessField = document.querySelector(''.guessField'');

        let guessCount = 1;
        let resetButton;
        function checkGuess() {
            const userGuess = Number(guessField.value);
            if (guessCount === 1) {
                guesses.textContent = ''Previous guesses: '';
            }
            guesses.textContent += `${userGuess} `;

            if (userGuess === randomNumber) {
                lastResult.textContent = ''Congratulations! You got it right!'';
                lastResult.style.backgroundColor = ''green'';
                lowOrHi.textContent = '''';
                setGameOver();
            } else if (guessCount === 10) {
                lastResult.textContent = ''!!!GAME OVER!!!'';
                lowOrHi.textContent = '''';
                setGameOver();
            } else {
                lastResult.textContent = ''Wrong!'';
                lastResult.style.backgroundColor = ''red'';
                if (userGuess < randomNumber) {
                    lowOrHi.textContent = ''Last guess was too low!'';
                } else if (userGuess > randomNumber) {
                    lowOrHi.textContent = ''Last guess was too high!'';
                }
            }

            guessCount++;
            guessField.value = '''';
            guessField.focus();
        }
        checkGuess();
        guessSubmit.addEventListener(''click'', checkGuess);
    </script>
</body>
</html>
'),(5,0,'index.html','<!DOCTYPE html>
<html lang="en-us">
  <head>
    <meta charset="utf-8">

    <title>Number guessing game</title>

    <style>
      html {
        font-family: sans-serif;
      }

      body {
        width: 50%;
        max-width: 800px;
        min-width: 480px;
        margin: 0 auto;
      }

      .form input[type="number"] {
        width: 200px;
      }

      .lastResult {
        color: white;
        padding: 3px;
      }
    </style>
  </head>

  <body>
    <h1>Number guessing game</h1>

    <p>We have selected a random number between 1 and 100. See if you can guess it in 10 turns or fewer. We''ll tell you if your guess was too high or too low.</p>

    <div class="form">
      <label for="guessField">Enter a guess: </label>
      <input type="number" min="1" max="100" required id="guessField" class="guessField">
      <input type="submit" value="Submit guess" class="guessSubmit">
    </div>

    <div class="resultParas">
      <p class="guesses"></p>
      <p class="lastResult"></p>
      <p class="lowOrHi"></p>
    </div>

    <script>
      let randomNumber = Math.floor(Math.random() * 100) + 1;
      const guesses = document.querySelector(''.guesses'');
      const lastResult = document.querySelector(''.lastResult'');
      const lowOrHi = document.querySelector(''.lowOrHi'');
      const guessSubmit = document.querySelector(''.guessSubmit'');
      const guessField = document.querySelector(''.guessField'');
      let guessCount = 1;
      let resetButton;

      function checkGuess() {
        const userGuess = Number(guessField.value);
        if (guessCount === 1) {
          guesses.textContent = ''Previous guesses: '';
        }

        guesses.textContent += userGuess + '' '';

        if (userGuess === randomNumber) {
          lastResult.textContent = ''Congratulations! You got it right!'';
          lastResult.style.backgroundColor = ''green'';
          lowOrHi.textContent = '''';
          setGameOver();
        } else if (guessCount === 10) {
          lastResult.textContent = ''!!!GAME OVER!!!'';
          lowOrHi.textContent = '''';
          setGameOver();
        } else {
          lastResult.textContent = ''Wrong!'';
          lastResult.style.backgroundColor = ''red'';
          if(userGuess < randomNumber) {
            lowOrHi.textContent = ''Last guess was too low!'' ;
          } else if(userGuess > randomNumber) {
            lowOrHi.textContent = ''Last guess was too high!'';
          }
        }

        guessCount++;
        guessField.value = '''';
        guessField.focus();
      }

      guessSubmit.addEventListener(''click'', checkGuess);

      function setGameOver() {
        guessField.disabled = true;
        guessSubmit.disabled = true;
        resetButton = document.createElement(''button'');
        resetButton.textContent = ''Start new game'';
        document.body.appendChild(resetButton);
        resetButton.addEventListener(''click'', resetGame);
      }

      function resetGame() {
        guessCount = 1;
        const resetParas = document.querySelectorAll(''.resultParas p'');
        for (const resetPara of resetParas) {
          resetPara.textContent = '''';
        }

        resetButton.parentNode.removeChild(resetButton);
        guessField.disabled = false;
        guessSubmit.disabled = false;
        guessField.value = '''';
        guessField.focus();
        lastResult.style.backgroundColor = ''white'';
        randomNumber = Math.floor(Math.random() * 100) + 1;
      }
    </script>
  </body>
</html>');