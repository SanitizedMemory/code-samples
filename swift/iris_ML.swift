//
//  main.swift
//  dimRed
//
//  Created by Michael Francis on 2020-11-08.
//

import Foundation
import TensorFlow
import PythonKit

let random = Python.import("random")
let datasets = Python.import("sklearn.datasets")
let plt = Python.import("matplotlib.pyplot")

let iris = datasets.load_iris()
let x = [[Float]](iris["data"].astype("float32"))!
let y = [Int](iris["target"])!

/*
 
 [
     [x, x, x, x],
     [x, x, x, x],
     [x, x, x, x]
 ]
 
 (x1, y1, z) => (x2, y2)
  
 (0, 0, 0)

 */

var weight = Tensor<Float>(randomUniform: [4, 2])
let xTensor = Tensor<Float>(x.map { Tensor($0) })

@differentiable
func lossFn(coords1: Tensor<Float>, coords2: Tensor<Float>, weight: Tensor<Float>) -> Tensor<Float> {
    //print("COORDS1")
//    print(coords1)
    //print("COORDS2")
//    print(coords2)
    let pc1 = matmul(coords1, weight)
    let pc2 = matmul(coords2, weight)
    //print("WEIGHT-FN")
//    print(weight)
    let guessedDist = sqrt(pow(pc1 - pc2, 2).sum(squeezingAxes: 1))
    let realDist = sqrt(pow(coords1 - coords2, 2).sum(squeezingAxes: 1))
    print("MSE")
    //print(pow(guessedDist - realDist, 2).mean())
    return pow(guessedDist - realDist, 2).mean()
}

func newMinibatch(batchSize: Int = 512) -> (Tensor<Float>, Tensor<Float>) {
    var batch1: [Tensor<Float>] = []
    var batch2: [Tensor<Float>] = []
    for _ in 1...batchSize {
        let idxs = [Int](random.sample(Python.range(x.count), 2))!
        batch1.append(Tensor(x[idxs[0]]))
        batch2.append(Tensor(x[idxs[1]]))
    }
    return (Tensor(batch1), Tensor(batch2))
}

for _ in 1...100 {
    let (a, b) = newMinibatch()
    let (loss, grad) = TensorFlow.valueWithGradient(at: weight) { weight in
        return lossFn(coords1: a, coords2: b, weight: weight)
    }
    print(loss)
    print("WEIGHT")
    print(weight)
    weight += grad * -0.05
    print("GRAD")
    print(grad)
    print("NEW WEIGHT")
    print(weight)
    print(grad * -0.05)
}

let yhat = [[Float]](matmul(xTensor, weight).makeNumpyArray())!
plt.scatter(yhat.map { $0[0] }, yhat.map { $0[1] }, c: y.map { ["red", "green", "blue"][$0] })
plt.show()

