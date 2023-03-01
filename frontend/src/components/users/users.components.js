import React, {useState, useEffect} from 'react';

import axios from "axios";

import {Button, Form, Container, Modal } from 'react-bootstrap'

import User from './single-user.component';
const BASE_URL = "http://localhost:8080"
const token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InNhbnRvc2gucmF5QHByb2RvLmluIiwiRmlyc3RfbmFtZSI6InNhbnRvc2giLCJMYXN0X25hbWUiOiJyYXkiLCJVaWQiOiI2MzYzN2EwYWYwMDBjZGNiZjczOWQ4YTUiLCJleHAiOjE2Nzc3NDU1NjF9._UlNbJ7Kt6tINkI33IXBhminiTgw4GU0cuWDrQ8VVQU"
const Users = () => {

    const [Users, setUsers] = useState([])
    const [refreshData, setRefreshData] = useState(false)

    const [changeUser, setChangeUser] = useState({"change": false, "id": 0})
    const [changeWaiter, setChangeWaiter] = useState({"change": false, "id": 0})
    const [newWaiterName, setNewWaiterName] = useState("")

    const [addNewUser, setAddNewUser] = useState(false)
    const [newUser, setNewUser] = useState({"first_name": "", "last_name": "", "email": "", "phone": ""})

    //gets run at initial loadup
    useEffect(() => {
        getAllUsers(); 
    }, [])

    //refreshes the page
    if(refreshData){
        setRefreshData(false);
        getAllUsers();
    }

    return (
        <div>
            
            {/* add new User button */}
            <Container>
                <Button onClick={() => setAddNewUser(true)}>Add new User</Button>
            </Container>

            {/* list all current Users */}
            <Container>
                {Users != null && Users.map((user, i) => (
                    <User UserData={user} deleteSingleUser={deleteSingleUser} setChangeWaiter={setChangeWaiter} setChangeUser={setChangeUser}/>
                ))}
            </Container>
            
            {/* popup for adding a new User */}
            <Modal show={addNewUser} onHide={() => setAddNewUser(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title>Add User</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form.Group>
                        <Form.Label >first_name</Form.Label>
                        <Form.Control onChange={(event) => {newUser.dish = event.target.value}}/>
                        <Form.Label>last_name</Form.Label>
                        <Form.Control onChange={(event) => {newUser.server = event.target.value}}/>
                        <Form.Label >email</Form.Label>
                        <Form.Control onChange={(event) => {newUser.table = event.target.value}}/>
                        <Form.Label >phone</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newUser.price = event.target.value}}/>
                    </Form.Group>
                    <Button onClick={() => addSingleUser()}>Add</Button>
                    <Button onClick={() => setAddNewUser(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>
            
            {/* popup for changing a waiter */}
            <Modal show={changeWaiter.change} onHide={() => setChangeWaiter({"change": false, "id": 0})} centered>
                <Modal.Header closeButton>
                    <Modal.Title>Change first_name</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form.Group>
                        <Form.Label >new User</Form.Label>
                        <Form.Control onChange={(event) => {setNewWaiterName(event.target.value)}}/>
                    </Form.Group>
                    <Button onClick={() => changeWaiterForUser()}>Change</Button>
                    <Button onClick={() => setChangeWaiter({"change": false, "id": 0})}>Cancel</Button>
                </Modal.Body>
            </Modal>

            {/* popup for changing an User */}
            <Modal show={changeUser.change} onHide={() => setChangeUser({"change": false, "id": 0})} centered>
                <Modal.Header closeButton>
                    <Modal.Title>Change User</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form.Group>
                        <Form.Label >first_name</Form.Label>
                        <Form.Control onChange={(event) => {newUser.dish = event.target.value}}/>
                        <Form.Label>last_name</Form.Label>
                        <Form.Control onChange={(event) => {newUser.server = event.target.value}}/>
                        <Form.Label >email</Form.Label>
                        <Form.Control onChange={(event) => {newUser.table = event.target.value}}/>
                        <Form.Label >phone</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newUser.price = parseFloat(event.target.value)}}/>
                    </Form.Group>
                    <Button onClick={() => changeSingleUser()}>Change</Button>
                    <Button onClick={() => setChangeUser({"change": false, "id": 0})}>Cancel</Button>
                </Modal.Body>
            </Modal>
        </div>
        
    );

    //changes the waiter
    function changeWaiterForUser(){
        changeWaiter.change = false
        var url = BASE_URL+"/waiter/update/" + changeWaiter.id
        axios.put(url, {
            "first_name": newWaiterName
        }).then(response => {
            console.log(response.status)
            if(response.status === 200){
                setRefreshData(true)
            }
        })
        
    }

    //changes the User
    function changeSingleUser(){
        changeUser.change = false;
        var url = BASE_URL+"/user/" + changeUser.id
        axios.put(url, newUser)
            .then(response => {
            if(response.status === 200){
                setRefreshData(true)
            }
        })
    }

    //creates a new User
    function addSingleUser(){
        setAddNewUser(false)
        var url = BASE_URL+"/user"
        axios.post(url, {
            "first_name": newUser.server,
            "last_name": newUser.dish,
            "email": newUser.table,
            "phone": parseFloat(newUser.price)
        }).then(response => {
            if(response.status === 200){
                setRefreshData(true)
            }
        })
    }

    //gets all the Users
    function getAllUsers(){
        // console.log("getAllUsers")
        var url = BASE_URL+"/users"
        const headers = {
                "Content-Type": "application/json",
                Accept: 'application/json',
                "Access-Control-Allow-Origin": true,
                token: token,
          };
        axios.get(url,{headers}).then(response => {
            if(response.status == 200){
                setUsers(response.data.data)
            }
        }).catch((error)=>{
            console.log("error",error)
        })
    }

    //deletes a single User
    function deleteSingleUser(id){
        var url = BASE_URL+"/user/" + id
        axios.delete(url, {

        }).then(response => {
            if(response.status === 200){
                setRefreshData(true)
            }
        })
    }

}

export default Users