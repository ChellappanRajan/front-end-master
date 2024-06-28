#Promise
    Promise is a object which will eventually complete or failure of an asynchronus opertion.

    Promise is a returned object which you attach callbacks, instead of passing call backs.

    function getLogInInformation(success,failure){
        function getUserDetails(details){
            function checkTokenExpiration(){
                if(notTokenExpired){
                    success(data);
                }else{
                    failure()
                }
            }
        }
    }

    const getLogInformation = //Promise();
    