// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

function calculatePressed() {
  // This function takes a user-given number and calculates the answer based on a pattern
  // Input through Textfields
  const layer = parseFloat(document.getElementById("layer").value)
  let addby = 3
  let answer = 0

  // Process
  for (let counter = 0; counter < layer; counter++) {
    answer += addby
    if (answer >= 6) {
      addby++
    }
  }

  // Output
  document.getElementById("answer").innerHTML =
    "Layer " + layer + " has " + answer + " dots."
}
