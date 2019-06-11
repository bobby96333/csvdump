# csvdump
dump csv file from remote data of mysql db

# Example

   > csvdump 
   > -e="select * from tb1 where id='1''"
   > -dsn="mysql://root:password@tcp(127.0.0.1:3006)/dbname" > output.csv
   
   
