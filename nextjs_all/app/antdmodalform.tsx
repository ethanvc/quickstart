'use client';
import { Button, Modal, Input } from 'antd';
import {useState} from "react";

let refreshCount = 0;
export function AntdModalForm(){
    const [openDlg, setOpenDlg] = useState(false);
    const [formData, setFormData] =
        useState({input1: "", input2:""});
    const onInputChange = (e :{target:{name:string, value:string}})=>{
        setFormData({...formData, [e.target.name]: e.target.value});
    };
    refreshCount++;
    return (
        <div data-id="nihao">
            <Button type="primary" onClick={()=>setOpenDlg(true)}>Open Modal</Button>
            <Modal title="Basic Modal" open={openDlg} onOk={()=>setOpenDlg(false)} onCancel={()=>setOpenDlg(false)}>
                <Input name="input1" placeholder="Basic usage" onChange={onInputChange}/>
                <Input name="input2" placeholder="Basic usage" onChange={onInputChange}/>
                <label>refresh count:{refreshCount}</label>
            </Modal>
        </div>
    );
}