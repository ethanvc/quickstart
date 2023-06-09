'use client';
import {useState} from "react";
import { Modal, Button, Checkbox, Label, TextInput} from 'flowbite-react';

let refreshCount = 0;
export function VerifyFlowbiteModal(){
    const [formData, setFormData] =
        useState({name:"", hobby: ""});
    const onInputChange = (e :{target:{name:string, value:string}})=>{
      setFormData({...formData, [e.target.name]: e.target.value});
    };
    refreshCount++;
    return (
        <div data-xx="nihao">
            <Modal show={true}>
            <Modal.Header>登录</Modal.Header>
            <Modal.Body>
                <TextInput name="name" onChange={onInputChange}></TextInput>
                <TextInput name="hobby" onChange={onInputChange}></TextInput>
                <Label>refresh count: {refreshCount}, data:{JSON.stringify(formData)}</Label>
            </Modal.Body>
        </Modal>
        </div>
    );

}