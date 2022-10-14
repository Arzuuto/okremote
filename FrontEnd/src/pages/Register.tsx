import { Flex, Text, Input, InputGroup, VStack, InputRightElement, useToast, Button, Link } from '@chakra-ui/react'
import { useState } from 'react'
import TopBar from "../components/Bar/TopBarThree"
import { useNavigate } from "react-router-dom";


export default function Register() {
    const toast = useToast()
    const navigate = useNavigate();
    const [Email, setEmail] = useState('')
    const [Password, setPassword] = useState('')
    const [ConfirmePass, setConfirmePass] = useState('')
    const [name, setName] = useState('')
    const [prenom, setPrenom] = useState('')
    const [show, setShow] = useState(false)
    const ShowOrHide = () => setShow(!show)

    const handleClick = async () => {
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, prenom, email: Email, password: Password }),
        };
        console.log(JSON.stringify({ email: Email, password: Password }))
        fetch('http://localhost:8080/register', requestOptions)
            .then((data) => {
                if (data.status === 200) {
                    navigate("/homeConnecte");
                    toast({
                        title: "Bonjour !",
                        status: "success",
                        duration: 5000,
                        isClosable: true,
                    });
                } else {
                    toast({
                        title: "une erreur sur a creation",
                        status: "error",
                        duration: 5000,
                        isClosable: true,
                    });
                }
            })
            .catch((error) => {
                console.log(error)
                toast({
                    title: "error",
                    status: "error",
                    duration: 2000,
                    isClosable: true,
                });
            })
    }

    const btn = (Email === "" || Password === "" || ConfirmePass === "" || name === "" || prenom === "") || ConfirmePass !== Password ?
        <Button disabled onClick={handleClick} width='100%' colorScheme="blue" borderRadius="40">Créer un compte</Button> :
        <Button onClick={handleClick} width='100%' colorScheme="blue" borderRadius="40">Créer un compte</Button>

    return (
        <>
            <TopBar />
            <Flex justify="center">
                <Flex marginTop="8%" w="35%" h="550px" justify="center" align="center" border="1px" shadow="2xl" borderRadius="40">
                    <VStack align="start" w="90%" maxW="400px" spacing={4}>
                        <Text fontSize={32} marginBottom="2" fontWeight="bold">Créer un compte</Text>
                        <Input onChange={e => setName(e.target.value)} placeholder='Name' width="100%" borderRadius="40" borderColor="black" />
                        <Input onChange={e => setPrenom(e.target.value)} placeholder='Prenom' width="100%" borderRadius="40" borderColor="black" />
                        <Input onChange={e => setEmail(e.target.value)} placeholder='Email' width="100%" borderRadius="40" borderColor="black" />
                        <InputGroup size='md'>
                            <Input
                                pr='4.5rem'
                                type={show ? 'text' : 'password'}
                                placeholder='Enter password'
                                onChange={e => setPassword(e.target.value)} borderRadius="40" borderColor="black" />
                            <InputRightElement width='4.5rem'>
                                <Button h='1.75rem' size='sm' onClick={ShowOrHide}>
                                    {show ? 'Hide' : 'Show'}
                                </Button>
                            </InputRightElement>
                        </InputGroup>
                        <InputGroup size='md'>
                            <Input
                                pr='4.5rem'
                                type={show ? 'text' : 'password'}
                                placeholder='Enter password'
                                onChange={e => setConfirmePass(e.target.value)} borderRadius="40" borderColor="black" />
                            <InputRightElement width='4.5rem'>
                                <Button h='1.75rem' size='sm' onClick={ShowOrHide}>
                                    {show ? 'Hide' : 'Show'}
                                </Button>
                            </InputRightElement>
                        </InputGroup>
                        {btn}
                        <Link href="/Login" width='100%'>
                            <Button width='100%' colorScheme="blue" borderRadius="40">Deja un compte ?</Button>
                        </Link>
                    </VStack>
                </Flex>
            </Flex>
        </>
    )
}