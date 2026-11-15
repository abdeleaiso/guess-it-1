# guess-it

## Description
program that given a number as standard input, prints out a range in which the next number provided should be.

## Project Structure
```bash
.
|____guess-it-dockerized
| |____index.html
| |____index.js
| |____ai
| | |____huge-range
| | |____mse
| | |____median
| | |____linear-regr
| | |____correlation-coef
| | |____nic
| | |____average
| | |____big-range
| |____README.md
| |____package.json
| |____data_sets
```
## Requirement
**Go 1.24.4** or newer.

## Set-up
```bash
$ git clone <guess-it-1-repo>
$ cd guess-it-1
```
## Usage
To test the student program, these commands should be ran to have the
dependencies needed and to start the webpage on the port 3000:

```console
docker compose up
```

After opening your browser of preference in the port
[3000](http://localhost:3000/), if you try clicking on any of the `Test Data`
buttons, you will notice that in the Dev Tool/ Console there is a message which
tells you that you need another guesser besides the student.

Adding a guesser is simple. You need to add in the URL a guesser, in other
words, the name of one of the files present in the `ai/` folder:

```console
?guesser=<name_of_guesser>
```

For example:

```console
?guesser=big-range
```

After that, choose which of the random data set to test. After that you can
wait for the program to test all of the values (boooooring), or you can click
`Quick` to skip the waiting and be presented with the results.

Since the website uses big data sets, we advise you to clear the displays
clicking on the `Clean` button after each test.

## Feature
**Statistical analysis**
- Mean calculation: Calculates the average of all previously received numbers.
- Standard deviation: Measures how spread out the input values are.
- Variance calculation: Measures the dispersion of the dataset.
- Data tracking: Maintains previously entered values to update statistical calculations.

**Prediction**
- Range prediction: Estimates the lower and upper boundaries of the next expected number.
- Dynamic prediction: Updates the predicted range whenever a new number is received.
- Adaptive calculations: Uses accumulated observations to improve its estimates.
- Configurable prediction strategy: Allows different statistical approaches to be tested, such as mean ± standard deviation.
- Outlier awareness: Identifies unusually high or low observations that may affect predictions.

**Performance**
- Continuous input: Reads numbers from standard input until EOF.
- Fast calculations: Uses efficient statistical calculations to avoid unnecessary processing.
- Lightweight implementation: Can run without external statistical libraries.
- Automated execution: Supports execution through the required script.sh file.

## Limitations
**Statistical limitations**
- Unpredictable data: Random data may have no detectable pattern, making accurate predictions difficult.
- Small datasets: Predictions based on a few observations may be unreliable.
- Outliers: Extreme values can significantly affect the mean and standard deviation.
- Assumption of normality: Some prediction methods assume that observations follow an approximately normal distribution.
- No guaranteed accuracy: The program estimates a possible range but cannot guarantee that the next value will fall inside it.

**Prediction limitations**
- No perfect prediction: The program cannot know the next number in advance.
- Changing distributions: If the data distribution changes, previous observations may become less useful.
- Trend detection: Statistical methods based only on mean and standard deviation might fail to recognize increasing or decreasing trends.
- Range-size trade-off: Wider ranges have a higher chance of containing the next value but receive a lower score when correct.

**Technical limitations**
- Sequential processing: The program processes observations one at a time.
- Memory usage: Storing every observation requires memory proportional to the dataset size, unless using online calculations.
- Input format: The program expects valid numerical input.
- Environment dependency: Execution requires a compatible runtime and the provided testing environment.
## Author
- <a href="https://github.com/abdeleaiso"/>Me</a>

## Licence
This project is licenced under the MIT. See the <a href="https://github.com/abdeleaiso/guess-it-1/blob/main/LICENSE"/>LICENCE</a> file for details