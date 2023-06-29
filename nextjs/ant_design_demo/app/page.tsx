'use client';
import { Button, Modal, Input } from 'antd';
import {useState} from "react";

export default function Home() {
  return (
      <ModalDemo/>
  )
}

function ModalDemo(){
  const [openDlg, setOpenDlg] = useState(false);
  const [formData, setFormData] = useState({input1: ""});
  const onInputChange = (e)=>{
    setFormData({...formData, [e.target.name]: e.target.value});
  };
  return (
    <div>
      <Button type="primary" onClick={()=>setOpenDlg(true)}>Open Modal</Button>
      <Modal title="Basic Modal" open={openDlg} onOk={()=>setOpenDlg(false)} onCancel={()=>setOpenDlg(false)}>
          <Input name="input1" placeholder="Basic usage" onChange={onInputChange}/>
      </Modal>
    </div>
  );
}
