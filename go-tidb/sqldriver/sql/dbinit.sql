USE test;
DROP TABLE IF EXISTS employee;

CREATE TABLE employee (
    `EmpId` INT,
    `EmpFName` VARCHAR(255),
    `EmpLname` VARCHAR(255),
    `Salary`   FLOAT,
    PRIMARY KEY (`EmpId`)
);

INSERT INTO  employee VALUES 
(1, "John", "Wick", 50000),
(2, "Miles", "Morales", 60000),
(3,  "Peter", "Parker", 70000);