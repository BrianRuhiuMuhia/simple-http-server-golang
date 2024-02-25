a simple http server that handles only get requests implemented in golang,it has three routes /,/getall,/getone/id,the / 
route display a html page,the /getall route display data from the data.json file it basically reads the file and displays
the content,the /getone/id takes the data from the data.json and searches for a unique id of the json object then displays
it on the page
