type Result = 'pass' | 'fail'

function verify(result: Result) {
  console.log('typeSctipt')
  if (result === 'pass') {
    console.log('Passed')
  } else {
    console.log("Failed")
  }
}