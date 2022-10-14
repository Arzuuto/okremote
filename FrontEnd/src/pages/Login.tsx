import { useState } from 'react'
import { Flex, Text, Input, InputGroup, VStack, InputRightElement, useToast, Button, Link } from '@chakra-ui/react'
import TopBar from "../components/Bar/TopBarThree"
import { useNavigate } from "react-router-dom";

export default function Login() {
    const toast = useToast()
    const navigate = useNavigate();
    const [show, setShow] = useState(false)
    const ShowOrHide = () => setShow(!show)
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const handleClick = async () => {
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password }),
        };
        fetch('http://localhost:8080/login', requestOptions)
            .then((data) => {
                if (data.status === 200) {
                    navigate("/homeConnecte");
                    toast({
                        title: "Bonjour !",
                        status: "success",
                        duration: 5000,
                        isClosable: true,
                    });
                }else {
                    toast({
                        title: "Votre compte n'existe pas",
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



    const btn = (email === "" || password === "") ?
        <Button disabled onClick={handleClick} width='100%' colorScheme="blue" borderRadius="40">Connexion</Button> :
        <Button onClick={handleClick} width='100%' colorScheme="blue" borderRadius="40">Connexion</Button>

    return (
        <>
            <TopBar />
            <Flex justify="center" >
                <Flex marginTop="8%" w="35%" h="400px" justify="center" align="center" border="1px" shadow="2xl" borderRadius="40" >
                    <VStack align="start" w="90%" maxW="400px" spacing={4}>
                        <Text fontWeight="bold" fontSize={32} marginBottom="2">Connexion</Text>
                        <Input borderColor="black" onChange={e => setEmail(e.target.value)} placeholder='Email' width="100%" borderRadius="40" />
                        <InputGroup size='md'>
                            <Input
                                pr='4.5rem'
                                type={show ? 'text' : 'password'}
                                placeholder='Enter password'
                                onChange={e => setPassword(e.target.value)} borderRadius="40" borderColor="black" />
                            <InputRightElement width='4.5rem'>
                                <Button h='1.75rem' size='sm' onClick={ShowOrHide} borderRadius="40">
                                    {show ? 'Hide' : 'Show'}
                                </Button>
                            </InputRightElement>
                        </InputGroup>
                        {btn}
                        <Link href="/register" width='100%'>
                            <Button width='100%' colorScheme="blue" borderRadius="40">Pas de compte ?</Button>
                        </Link>
                    </VStack>
                </Flex>
            </Flex>
        </>
    )
}