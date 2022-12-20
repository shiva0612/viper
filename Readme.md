using Viper tool for reading the any config file (json xml yaml)
and marshalling it into respective config struct in golang 
for the application to use

following thigs can be achieved using viper:
1.  reading from config file (unmarshalling to struct (or) directly getting values "key.subkey.subsubkey.1")
2.  reading from raw string 
3.  watching for config file changes and running certain function
4.  override config file with env, flag
5.  use multiple viper for reading multiple config files
6.  once config file is loaded -> configs are overrided by env,falgs -> u can write the config used to run the server to a file "configUsed.json"