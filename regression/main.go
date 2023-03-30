/*import (
	"fmt"
	"log"
	"os"
)

// Open the CSV file.

func main() {
	advertFile, err := os.Open("C:/Users/tirdesh/OneDrive/Documents/fullstackdev_course/dataset/student_scores.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	// Create a dataframe from the CSV file.
	advertDF := dataframe.ReadCSV(advertFile)
	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	advertSummary := advertDF.Describe()
	// Output the summary statistics to stdout.
	fmt.Println(advertSummary)
}*/


/*Linear Regression in Golang for predicting
whether or not a user on social media will purchase a product through ads,
given the age, salary etc of the user */

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"

	//for plotting
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	//for regression
	"github.com/sajari/regression"
)

func main() {
	// Step 1: To open the desired dataframe
	netads, err := os.Open("C:/Users/tirdesh/OneDrive/Documents/fullstackdev_course/dataset/student_scores.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer netads.Close()

	// Step 2: Creating a dataframe from the CSV file
	/* Like we use Pandas in Python to create a dataframe, there seems
	   to be a method called dataframe.ReadCSV() that does the same from a .csv file */
	netdf := dataframe.ReadCSV(netads)

	// Step 3: Using the Describe() method to get stats like in Python
	netSummary := netdf.Describe()
	fmt.Println(netSummary)

	//Step 4: Open the dataset file for visualization and better understanding
	/*f, err := os.Open("C:/Users/hi/OneDrive/Documents/practice/Social_Network_Ads.csv")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()*/

	// step 4.1: creating a histogram for each of the cols in the dataset
	for _, colName := range netdf.Names() {

		//using a plotter.Values to fill 'plots' with the columns of the dataframe
		plots := make(plotter.Values, netdf.Nrow())
		for i, floater := range netdf.Col(colName).Float() {
			plots[i] = floater
		}

		//creating a plot
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)
		// Create a histogram of our values drawn
		// from the standard normal.
		h, err := plotter.NewHist(plots, 15)
		if err != nil {
			log.Fatal(err)
		}
		//normalizing the histogram
		h.Normalize(1)
		//adding the histogram to the plot
		p.Add(h)

		//saving it as an image
		if err := p.Save(3*vg.Inch, 3*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}

	// Step 5 Choosing the independent variable
	/* I have tried working with only one visualization after cloning the plot repo
	as the main aim of this assignment is to perform Linear Regression */

	// Opening the netdf file again
	f, err := os.Open("C:/Users/hi/OneDrive/Documents/practice/Social_Network_Ads.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//creating the dataframe
	netdf2 := dataframe.ReadCSV(f)

	// picking the target feature
	yTar := netdf.Col("Purchased").Float()

	//creating a scatter plot for all of the features
	for _, colName := range netdf.Names() {
		//declaring new variable to store values for plotting
		pts := make(plotter.XYs, netdf.Nrow())

		//Filling pts variable with the data
		//setting an iteratror that will pass through all cols and store all values
		// a separate one for the target variable too
		for i, floater := range netdf.Col(colName).Float() {
			pts[i].X = floater
			pts[i].Y = yTar[i]
		}

		//creating the plot
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = 'y'
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)
		// Save the plot to a PNG file.
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err !=
			nil {
			log.Fatal(err)
		}
		//not plotting for UserID as it does not bring any difference to the evaluations or generalizations

	}

	// Step 6: Splitting dataset into training and test datasets
	//  - calculate number of elements in each set
	training := (4 * netdf.Nrow()) / 5
	test := netdf.Nrow() / 5
	if training+test < netdf.Nrow() {
		training++
	}

	//creating index for the train and test subsets
	trainid := make([]int, training)
	testid := make([]int, test)

	//enumerate the training index
	//this will just set the index of whichever serial value it is to that serial
	// so if it is at the 3rd value the index will be set to 3 and so on
	for i := 0; i < training; i++ {
		trainid[i] = i
	}

	//similarly enumerating test index
	for i := 0; i < test; i++ {
		testid[i] = training + i
	}

	//create the datasets
	traindf := netdf.Subset(trainid)
	testdf := netdf.Subset(testid)

	//I need to create a 'map' that will be used in writing the data to the files
	setmap := map[int]dataframe.DataFrame{
		0: traindf,
		1: testdf,
	}

	//creating respective files
	for id, setName := range []string{"training.csv", "test.csv"} {

		//saving the filtered dataset
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}
		//creating a buffered writer in order to write a specific set of values at a time
		w := bufio.NewWriter(f)

		//write the dataframe as a csv file
		if err := setmap[id].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
	// Step 7: Training the model
		// opening the training dataset file
		f1, err := os.Open("training.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer f1.Close()

		//creating a new csv reader for this file
		reader := csv.NewReader(f1)

		//reading all of the records
		reader.FieldsPerRecord = 4
		trainingdata, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		// in this case the target variable is Purchased that is y
		// creating struct to train the model using the regression module
		var r regression.Regression
		r.SetObserved("Purchased")
		r.SetVar(0, "Age")

		//looping for records and adding the training data to the regression vals
		for i, record := range trainingdata {
			//skipping header
			if i == 0 {
				continue
			}
			//parse the Purchased regression or the y measure
			yTarval, err := strconv.ParseFloat(record[3], 64)
			if err != nil {
				log.Fatal(err)
			}

			//parse the Age value
			ageval, err := strconv.ParseFloat(record[0], 64)
			if err != nil {
				log.Fatal(err)
			}

			//adding these points to the regression value
			r.Train(regression.DataPoint(yTarval, []float64{ageval}))
		}

		// fit the regression model
		r.Run()

		//outputting the model params
		fmt.Printf("Regression Formula: \n%v\n\n", r.Formula)
    
	// Step 8: Evaluating the model
	//opening the test dataset
	f, err = os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	   }
	   defer f.Close()
	   // creating a csv reader to read the test file 
	   reader = csv.NewReader(f)

	   // reading all the records 
	   reader.FieldsPerRecord = 4
	   testdata, err := reader.ReadAll()
	   if err != nil {
		log.Fatal(err)
	   }
	   
	   //looping the prediction for y and evaluating that prediction with MAE or mean absolute error as the loss function
	   var mAE float64
	   for i, record := range testdata {
		// skipping the first index
		if i == 0 {
		continue
		}
		// Parse the observed predictions for "Purchased" or "y".
		yObs, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
		log.Fatal(err)
		}
		// Parse the age value.
		ageval, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
		log.Fatal(err)
		}
		// Predict y with trained model.
 yPredicted, err := r.Predict([]float64{ageval})

		//add the mae
		mAE += math.Abs(yObs-yPredicted)/float64({ageageval})
	   }

	   //output MAE to standard out
	   fmt.Println("Mean Absolute Error= %0.2f\n\n", mAE)


	   /* Running this will give the error value. Smaller the error, larger the accuracy score.
	   In roder to understand the overall predictions, a plot can be prepared as well. */

	//Step 8: Visualizing the regression 
	   for i, floatVal := range advertDF.Col("Age").Float() {
		pts[i].X = floater
		pts[i].Y = yTar[i]
		ptsPred[i].X = floater
		ptsPred[i].Y = predict(floater)
	   }
	   // Create the plot.
	   p, err := plot.New()
	   if err != nil {
		log.Fatal(err)
	   }
	   p.X.Label.Text = "Age"
	   p.Y.Label.Text = "Purchased: Yes or No"
	   p.Add(plotter.NewGrid())
	   // Add the scatter plot points for the observations.
	   s, err := plotter.NewScatter(pts)
	   if err != nil {
		log.Fatal(err)
	   }
	   s.GlyphStyle.Radius = vg.Points(3)
	   // Add the line plot points for the predictions.
	   l, err := plotter.NewLine(ptsPred)
	   if err != nil {
	} 
}
