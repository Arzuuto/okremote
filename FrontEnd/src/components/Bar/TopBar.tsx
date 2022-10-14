import { Button, HStack, IconButton, InputGroup, InputLeftElement, Link, Input } from '@chakra-ui/react'
import { Box } from '@chakra-ui/react'
import { SearchIcon } from '@chakra-ui/icons'
import { AiFillHome } from "react-icons/ai";
import React from 'react';

export default function TopBarDisconnect() {
    const [value, setValue] = React.useState('')
    const handleChange = (event: { target: { value: React.SetStateAction<string>; }; }) => setValue(event.target.value)

    return (
        <HStack bg='gray.800' w='100%' p={4} color='white' justify='space-between'>
            <HStack>
                <Link href="home">
                    <IconButton
                        colorScheme='teal'
                        aria-label='Home'
                        icon={<AiFillHome />} />
                </Link>
            </HStack>
            <HStack>
                <Link href="home">
                    <Button
                        fontWeight='bold'
                        fontSize='50px'
                        colorScheme='gray.800'
                        aria-label='Logo'
                    >Remote OK
                    </Button>
                </Link>
            </HStack>
            <HStack spacing='24px'>
                <Box>
                    <InputGroup>
                        <Input
                            value={value}
                            onChange={handleChange}
                            placeholder='Search' />
                        <InputLeftElement children={<SearchIcon color='teal' />} />
                    </InputGroup>
                </Box>
                <Box>
                    <Link href="/register">
                        <Button colorScheme='teal' mr='4'>Créer un compte</Button>
                    </Link>
                    <Link href="login">
                        <Button colorScheme='teal'>Connexion</Button>
                    </Link>
                </Box>
            </HStack>
        </HStack>
    )
}