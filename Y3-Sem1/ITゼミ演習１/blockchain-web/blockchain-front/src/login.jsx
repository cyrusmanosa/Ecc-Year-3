import Box from '@mui/material/Box';
import InputAdornment from '@mui/material/InputAdornment';
import TextField from '@mui/material/TextField';
import AccountCircle from '@mui/icons-material/AccountCircle';
import { useState, useCallback } from 'react';
import axios from 'axios';

export default function InputWithIcon() {
    const [name, setName] = useState('')
    const [email, setEmail] = useState('')
    const [company_name, setCompanyName] = useState('')
    const [message, setMessage] = useState('')
    const [uploadedFiles, setUploadedFiles] = useState([]); // 新增状态以存储上传的文件

    const handleName = useCallback((e) => {
        setName(e.target.value);
    }, []);

    const handleEmail = useCallback((e) => {
        setEmail(e.target.value);
    }, []);

    const handleCompanyName = useCallback((e) => {
        setCompanyName(e.target.value);
    }, []);

    const handleMessage = useCallback((e) => {
        setMessage(e.target.value);
    }, []);

    const handleFileChange = (e) => { // 新增文件处理函数
        setUploadedFiles(e.target.files);
    };

    const handleSendData = async () => {
        const data = {
            name,
            email,
            company_name,
            message,
            files: uploadedFiles // 将文件添加到数据中
        };
        const formData = new FormData(); // 创建 FormData 对象
        for (const file of uploadedFiles) {
            formData.append('files', file); // 将每个文件添加到 FormData
        }
        formData.append('data', JSON.stringify(data)); // 将其他数据添加到 FormData
        await axios.put("http://localhost:8080/take", formData, {
            headers: {
                'Content-Type': 'multipart/form-data' // 设置请求头
            }
        });
    }

    return (
        <Box sx={{ '& > :not(style)': { m: 1 } }}>
            {/* Name */}
            <TextField
                id="Name"
                label="Name"
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <AccountCircle />
                        </InputAdornment>
                    ),
                }}
                variant="standard"
                value={name}
                onChange={handleName}
            />
            {/* Email */}
            <TextField
                id="Email"
                label="Email"
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <AccountCircle />
                        </InputAdornment>
                    ),
                }}
                variant="standard"
                value={email}
                onChange={handleEmail}
            />
            {/* CompanyName */}
            <TextField
                id="CompanyName"
                label="CompanyName"
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <AccountCircle />
                        </InputAdornment>
                    ),
                }}
                variant="standard"
                value={company_name}
                onChange={handleCompanyName}
            />
            {/* Message */}
            <TextField
                id="Message"
                label="Message"
                InputProps={{
                    startAdornment: (
                        <InputAdornment position="start">
                            <AccountCircle />
                        </InputAdornment>
                    ),
                }}
                variant="standard"
                value={message}
                onChange={handleMessage}
            />
            <input
                type="file"
                id="fileUpload"
                name="uploaded"
                multiple
                onChange={handleFileChange} // 更新文件处理函数
            />
            <button onClick={handleSendData}>Submit</button>
        </Box>
    );
}
