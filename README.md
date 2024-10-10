# MSDS 431 - Creating a Command-Line Application
This is document outlines the testing procedures and how to run this go application. The intended use of the executable is to convert housing data into a json file.

*Note* the Following appication will not convert all csv files into JSON files and will only work with housing data fomatted as menntioned in [this section](CSV Specifications)

**Assignment Scenario**
A data science consulting firm would like to have a utility for converting comma-delimited text (comma-separated values, csv) to JavaScript Object Notation (JSON).  The JSON lines file represents a common file format for populating databases and document stores.

We take the role of company data scientists, learning how to use Go on the command line.

Your job is to write a Go command-line application to read the comma-delimited text file for California housing data and create a JSON lines file for these same data as output. It would be most convenient for the command-line application to have arguments for the path to the input csv file and the path to the output JSON lines file. 

**Helpful Resources Used Creating this Application**

[Markdown Cheatsheet](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#links)

[JSON File Validator](https://jsonlint.com/)

[Overview of working with Structs and JSON files in Go](https://www.sohamkamani.com/golang/json/#structured-data-decoding-json-into-structs)

## CSV Specifications
The inputted .csv file must follow a specific format outlined below. Please reference [housesInput.csv](housesInput.csv) for the file that was used to test and design this application:
* The first row of data should be a header.
* The data should be organized into the following order:
  
  * Value - Median house value
  * Income - Median Income
  * age - Housing median age
  * rooms - Total Rooms
  * bedrooms - Total Bedrooms
  * pop - Population
  * hh - Households
  
  Any columns past hh will be ignored.
*All data outside csv the header should be a numerical value.

## Using the Application
The [HousingConversion.exe](HousingConversion.exe) file in this executable is intended for Windows machines. Download the executable and place the file in a location that you want to reference.

Open up coommand prompt anc change the directory using the ```cd``` function. Below is a scenario in which I place the executable on my desktop.

```
cd C:\Users\migus\OneDrive\Documents\Desktop
```
To run the executable, use the following command ```HousingConversion [InputFilePath] [OutputFilePath]```
* ```[InputFilePath]``` should be the .csv file name if if the csv is in the same directory as your executable _or_ in the entire filepath including the filename should be specified.

  Below are 2 inputs that would work as the InputFilePath:
  * ```"C:\Users\migus\OneDrive\Documents\Desktop\housesInput.csv"```
  * ```housesInput.csv```
    * _assuming the the csv is stored in the same working directory as the executable_
* ```[OutputFilePath]``` represents the generated JSON file name. Similar to [InputFilePath], if the absolute path of where you would like the file to be generated is not provided, the default behaviour will be to use the directory where the executable is stored

Puttint it all together, the follwing command would convert housing data from CSV to JSON assuming both input and output files are stored in the same location as the executable:

```HousingConversion housesInput.csv outputData.json```

Alternatively, this would achieve the same exact outcome:

```HousingConversion "C:\Users\migus\OneDrive\Documents\Desktop\housesInput.csv" "C:\Users\migus\OneDrive\Documents\Desktop\outputData.json\```

## Testing the Application
### Valid User Inputs
### Input/Output File Validation

