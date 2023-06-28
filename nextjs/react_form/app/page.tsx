'use client';
// import { ChakraProvider, Input, Button} from '@chakra-ui/react'
import {useState} from "react";
import { Button , TextInput} from 'flowbite-react';

export default function Home() {
  return (
      // <ChakraProvider>
          <RealHome/>
      // </ChakraProvider>
  )
}

function RealHome(){
    const [formData, setFormData] = useState({name:"", hobby: ""});
    return (
        <div>
            <label>name</label>
            <TextInput onChange={e=>setFormData({...formData, name: e.target.value})}></TextInput>
            <label>hobby</label>
            <TextInput onChange={e=>setFormData({...formData, hobby: e.target.value})}></TextInput>
            <Button onClick={()=>console.log(formData)}>Result</Button>
        </div>
    );
}
