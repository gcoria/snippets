//Filter regex expressions
function array_filter(arr) {
  return arr.filter(x => /'change_me'/.test(x)).length;
}

// Destructuring array of tuples
function destructure_tuples(data){
  return data.map(([prop1, prop2]) => (prop1 > 54 && prop2 > 7) ? 'Senior' : 'Open');
}
