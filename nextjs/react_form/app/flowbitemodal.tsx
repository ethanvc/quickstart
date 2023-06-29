import {useState} from "react";
import { Modal, Button, Checkbox, Label, TextInput} from 'flowbite-react';

export function VerifyFlowbiteModal(){
    const [formData, setFormData] =
        useState({name:"", hobby: ""});
    const onInputChange = (e)=>{
      setFormData({...formData, [e.target.name]: e.target.value});
    };
    return (
        <div>
            <Modal show={true}>
            <Modal.Header>登录</Modal.Header>
            <Modal.Body>
                <TextInput name="name" onChange={onInputChange}></TextInput>
                <TextInput name="hobby" onChange={onInputChange}></TextInput>
                <Label>{JSON.stringify(formData)}</Label>
            </Modal.Body>
        </Modal>
        </div>
    );

}