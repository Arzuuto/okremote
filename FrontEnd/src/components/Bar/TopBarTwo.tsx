import { Button, HStack, IconButton, Link, Box, Text } from '@chakra-ui/react'
import { AiFillHome } from "react-icons/ai";
import { useState, useEffect } from 'react';

export default function TopBarConnected() {
    const [totalReactPackages, setTotalReactPackages] = useState(null);

    const requestOptions2 = {
        method: "GET",
        headers: { "Content-Type": "application/json" }
    };

    fetch("http://localhost:8080/homeConnecte", requestOptions2)
        .then((data) => {
            console.log(data);
        })
    return (
        <HStack bg='gray.800' w='100%' p={4} color='white' justify='space-between'>
            <HStack>
                <Link href="/homeConnecte">
                    <IconButton
                        colorScheme='teal'
                        aria-label='Home'
                        icon={<AiFillHome />} />
                </Link>
            </HStack>
            <HStack>
                <Link href="/homeConnecte">
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
                    <Text></Text>
                </Box>
                <Box>
                    <Link href="/CreateAnnonce">
                        <Button colorScheme='teal' mr='4'>Créer une anonce</Button>
                    </Link>
                    <Link href="/home">
                        <Button colorScheme='teal'>Déconnexion</Button>
                    </Link>
                </Box>
            </HStack>
        </HStack>
    )
}