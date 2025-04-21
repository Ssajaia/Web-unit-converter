Unit Converter Web Application
Overview
This is a simple web-based unit converter that allows users to convert between different units of measurement including length, weight, and temperature. The application is built with:

Backend: Go (Golang)

Frontend: HTML templates with CSS styling

Architecture: Traditional server-side rendering

Features
Convert between various units:

Length: millimeters, centimeters, meters, kilometers, inches, feet, yards, miles

Weight: milligrams, grams, kilograms, ounces, pounds

Temperature: Celsius, Fahrenheit, Kelvin

Clean, intuitive interface

Responsive design that works on desktop and mobile devices

No database required - all conversions calculated in real-time

Installation
Prerequisites
Go 1.16 or higher

Git (optional)

Steps
Clone the repository:

bash
git clone https://github.com/yourusername/unit-converter.git
cd unit-converter
Run the application:

bash
go run main.go
Open your browser and visit:

http://localhost:8080
Project Structure
unit-converter/
├── main.go                # Main application entry point
├── static/
│   └── style.css          # CSS stylesheet
├── templates/
│   ├── base.html          # Base template
│   ├── length.html        # Length converter page
│   ├── weight.html        # Weight converter page
│   └── temperature.html   # Temperature converter page
└── README.md              # This file
Usage
Select the conversion type (Length, Weight, or Temperature) from the navigation bar

Enter the value you want to convert

Select the unit to convert from

Select the unit to convert to

Click "Convert" to see the result

Conversion Formulas
Length
All length conversions are done via meter as the base unit

Example: 1 foot = 0.3048 meters

Weight
All weight conversions are done via gram as the base unit

Example: 1 pound = 453.592 grams

Temperature
Special formulas for temperature conversions:

Celsius to Fahrenheit: °F = (°C × 9/5) + 32

Fahrenheit to Celsius: °C = (°F - 32) × 5/9

Celsius to Kelvin: K = °C + 273.15

Troubleshooting
If you encounter issues:

White screen:

Verify all template files exist in the templates directory

Check that templates have proper {{define "content"}} sections

Ensure the Go server is running without errors

Missing styles:

Confirm the static/style.css file exists

Check browser console for 404 errors

Conversion errors:

Verify you entered a valid number

Check server logs for any calculation errors

Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements.

https://roadmap.sh/projects/unit-converter
