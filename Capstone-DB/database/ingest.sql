CREATE TABLE `DIM_CUSTOMER`
(
    id int auto_increment,
    Customer_Name varchar(200) NOT NULL, 
    City varchar(150) NOT NULL, 
    PRIMARY KEY (`id`)
); 

LOAD DATA LOCAL INFILE '/input/sample.csv' 
INTO TABLE DIM_CUSTOMER 
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;
