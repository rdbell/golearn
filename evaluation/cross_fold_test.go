package evaluation

import (
	"testing"

	"github.com/rdbell/golearn/base"
	"github.com/rdbell/golearn/knn"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCrossFold(t *testing.T) {
	Convey("Cross Fold Evaluation", t, func() {
		iris, _ := base.ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
		cls := knn.NewKnnClassifier("euclidean", "linear", 2)
		cfs, _ := GenerateCrossFoldValidationConfusionMatrices(iris, cls, 5)
		Convey("Cross Fold Validation Confusion Matrices", func() {
			So(cfs, ShouldNotBeEmpty)
		})
		Convey("Cross Validated Metric", func() {
			mean, variance := GetCrossValidatedMetric(cfs, GetAccuracy)
			So(mean, ShouldBeBetween, 0.8, 1)
			So(variance, ShouldBeBetween, 0, 0.03)
		})
	})
}
