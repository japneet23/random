create database healthcare;
use healthcare;

CREATE TABLE Doctor2 (
doctor_id INTEGER NOT NULL primary key,
doctor_name varchar(20) NOT NULL,
speciality varchar(20) NOT NULL);

CREATE TABLE patient2(
patient_id INTEGER NOT NULL primary key,
patient_name VARCHAR(20) NOT NULL);

CREATE TABLE appointment8(
patient_id INTEGER NOT NULL,
doctor_id INTEGER NOT NULL,
starttime datetime not null ,
endtime datetime not null ,
Primary key(doctor_id,starttime),
FOREIGN KEY (doctor_id) REFERENCES Doctor2 (doctor_id),
FOREIGN KEY (patient_id) REFERENCES patient2 (patient_id));

insert into Doctor2
values(101,'bill','cardiology');
insert into Doctor2
values(102,'king','physician');
insert into Doctor2
values(103,'lucky','oncology');
insert into Doctor2
values(104,'Sam','neurology');
insert into Doctor2
values(105,'Victor','surgical');
insert into Doctor2
values(106,'Vlad','heamatology');

insert into Doctor2
values(107,'tripper','cardiology');
insert into Doctor2
values(108,'gonna','physician');
insert into Doctor2
values(109,'trip','oncology');
insert into Doctor2
values(110,'me','oncology');

insert into patient2
values(201,'Anil');
insert into patient2
values(202,'Ram');
insert into patient2
values(203,'Klu');
insert into patient2
values(204,'Klux');
insert into patient2
values(205,'Klan');
insert into patient2
values(206,'whew');

insert into patient2
values(207,'I');
insert into patient2
values(208,'am');
insert into patient2
values(209,'going');
insert into patient2
values(210,'to');
insert into patient2
values(211,'be');
insert into patient2
values(212,'godly');


insert into appointment2
values(201,101,'2021-01-02 17:30:00','2021-01-02 18:00:00');
insert into appointment2
values(202,102,'2021-01-02 18:00:00','2021-01-02 18:30:00');
insert into appointment2
values(203,104,'2021-01-02 09:30:00','2021-01-02 10:00:00');
insert into appointment2
values(204,101,'2021-01-02 10:00:00','2021-01-02 10:30:00');
insert into appointment2
values(205,106,'2021-01-03 11:30:00','2021-01-03 12:00:00');
insert into appointment2
values(206,103,'2021-01-02 12:00:00','2021-01-02 12:30:00');
insert into appointment2
values(207,105,'2021-01-03 14:00:00','2021-01-03 14:30:00');
insert into appointment2
values(208,104,'2021-01-02 16:00:00','2021-01-02 16:30:00');
insert into appointment2
values(209,101,'2021-01-02 08:00:00','2021-01-02 08:30:00');
insert into appointment2
values(210,102,'2021-01-02 07:30:00','2021-01-02 08:00:00');
insert into appointment2
values(211,103,'2021-01-09 09:30:00','2021-01-09 10:00:00');
insert into appointment2
values(212,101,'2021-01-02 17:00:00','2021-01-02 17:30:00');
insert into appointment2
values(212,109,'2021-01-07 09:00:00','2021-01-07 09:30:00');
insert into appointment2
values(212,107,'2021-01-05 10:00:00','2021-01-05 10:30:00');

select speciality, count(doctor_name)
from Doctor2
group by speciality;

select * from appointment
where patient_id=212;


SELECT doc.doctor_names,free_time=app.endtime
FROM (
    SELECT @lastEndTime as AvailStartTime, StartTime as AvailEndTime, @lastEndTime := EndTime
    FROM (SELECT StartTime, EndTime
          FROM Events
          WHERE Pers_ID = 101
            AND EndTime >= '2013-09-21 09:00'
            AND StartTime < '2013-09-21 22:00'
          ORDER BY StartTime) e
    JOIN (SELECT @lastEndTime := NULL) init) x
WHERE AvailEndTime > DATE_ADD(AvailStartTime, INTERVAL 30 MINUTE)