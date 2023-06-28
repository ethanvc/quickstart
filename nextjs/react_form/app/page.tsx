'use client';
// import { ChakraProvider, Input, Button} from '@chakra-ui/react'
import {useState} from "react";
import { Button , TextInput} from 'flowbite-react';
import ReactDOM from "react-dom";

export default function Home() {
  return (
      // <ChakraProvider>
          <RealHome/>
      // </ChakraProvider>
  )
}

function RealHome(){
    console.log("rerender");
    const [formData, setFormData] = useState({name:"", hobby: ""});
    const content = (
        <div>
            <label>name</label>
            <TextInput key="name" onChange={e=>setFormData({...formData, name: e.target.value})}></TextInput>
            <label>hobby</label>
            <TextInput key="hobby" onChange={e=>setFormData({...formData, hobby: e.target.value})}></TextInput>
            <Button onClick={()=>console.log(formData)}>Result</Button>
        </div>
    );
    return ReactDOM.createPortal(content, document.body);
}
