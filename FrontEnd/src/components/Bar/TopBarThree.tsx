import { Button, HStack, IconButton, Link, Box, Text, Flex } from '@chakra-ui/react'
import { AiFillHome } from "react-icons/ai";

export default function TopBarConnecting() {

    return (
        <HStack bg='gray.800' w='100%' p={4} color='white' spacing='37%'>
            <Flex justify='start'>
                <Link href="home">
                    <IconButton
                        colorScheme='teal'
                        aria-label='Home'
                        icon={<AiFillHome />} />
                </Link>
            </Flex>
            <Flex justify='center'>
                <Link href="home">
                    <Button
                        fontWeight='bold'
                        fontSize='50px'
                        colorScheme='gray.800'
                        aria-label='Logo'
                    >Remote OK
                    </Button>
                </Link>
            </Flex>
        </HStack>
    )
}