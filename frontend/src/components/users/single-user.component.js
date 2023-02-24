import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import {Button, Card, Row, Col} from 'react-bootstrap'
const User = ({UserData, setChangeWaiter, deleteSingleUser, setChangeUser}) => {
    return (
        <Card>
    <Row>
        <Col>First name:{ UserData !== undefined && UserData.first_name}</Col>
        <Col>last name:{ UserData !== undefined && UserData.last_name} </Col>
        <Col>Email:{ UserData !== undefined && UserData.email}</Col>
        <Col>Phone No: ${UserData !== undefined && UserData.phone}</Col>
        <Col><Button onClick={() => deleteSingleUser(UserData.id)}>delete User</Button></Col>
        <Col><Button onClick={() => changeWaiter()}>change waiter</Button></Col>
        <Col><Button onClick={() => changeUser()}>change User</Button></Col>
    </Row>
</Card>
    )
    function changeWaiter(){
        setChangeWaiter({
            "change": true,
            "id": UserData._id
        })
    }
    function changeUser(){
        setChangeUser({
            "change": true,
            "id": UserData._id
        })
    }
}
export default User