function doSomething() {
    return new Promise((resolve) => {
      setTimeout(() => {
        // Other things to do before completion of the promise
        console.log("Did something");
        // The fulfillment value of the promise
        resolve("https://example.com/");
      }, 200);
    });
  }
 
  
function C(){
    return new Promise((resolve) => {
    //   setTimeout(() => {
        // Other things to do before completion of the promise
        console.log("Did something C");
        // The fulfillment value of the promise
        resolve("https://example.com/");
    //   }, 200);
    });
  }
 
 
function B() {
    return new Promise((resolve) => {
     setTimeout(() => {
        // Other things to do before completion of the promise
        console.log("Did something B");
        // The fulfillment value of the promise
        resolve("https://example.com/");
      }, 200);
    });
  }

//Run parally
//If one rejected means it will abort other operations
  Promise.all([C(),B()]).then((r)=>{
    console.log(r);
  })
 
//   doSomething()
//   .then((result) => {
//    return result
//   })
//   .then((newResult) => {
//     throw new Error('Some Error');
//   })
//   .then((finalResult) => console.log(`Got the final result: ${finalResult}`))
//   .catch((e)=>{console.log(e)}).finally(()=>{
//     console.log('Finally Ran')
//   })
