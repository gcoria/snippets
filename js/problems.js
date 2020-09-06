// 3d array max hourglass
function hourglassSum(arr) {
    let max = -63;

    for (let i = 0; i < 4; i++) {
        for (let j = 0; j < 4; j++) {
            let sum = arr[i + 1][j + 1];
            for (let k = 0; k < 3; k++) {
                sum += arr[i][j + k];
                sum += arr[i + 2][j + k]
            }
            if (sum > max) max = sum;
        }
    }

    return max;
}

// filter with smiley faces
function countSmileys(arr) {
  return arr.filter(x => /^[:;][-~]?[)D]$/.test(x)).length;
}


//insert array 1 at index n of arr2 
function frankenSplice(arr1, arr2, n) {
  let newArray = arr2.slice(); // slice COPIA la referencia dentro del nuevo array
  newArray.splice(n, 0, ...arr1); // inserta en el indice n, elimina 0 elementos
  return newArray;
}


// Destructuring
// Arrow Functions: () => {} 
function openOrSenior(data){
  return data.map(([age, handicap]) => (age > 54 && handicap > 7) ? 'Senior' : 'Open');
}


// Change - Movies
function Clerk(name) {
  this.name = name;
  
  this.money = {
    25 : 0,
    50 : 0,
    100: 0 
  };
  
  this.sell = function(element, index, array) {
    this.money[element]++;

    switch (element) {
      case 25:
        return true;
      case 50:
        this.money[25]--;
        break;
      case 100:
        this.money[50] ? this.money[50]-- : this.money[25] -= 2;
        this.money[25]--;
        break;
    }
    return this.money[25] >= 0;
  };
}

function tickets(peopleInLine){
  var vasya = new Clerk("Vasya");
  return peopleInLine.every(vasya.sell.bind(vasya)) ? "YES" : "NO";
}
// ************ Change - Movies