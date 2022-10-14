import { Image, Badge, HStack, Button, Text, VStack } from '@chakra-ui/react';
import { useState, useEffect } from 'react';

type Property = {
  Contrat: string;
  Date: string;
  Image: string;
  Localisation: string;
  Name: string;
  Remuneration: string;
  Sector: string;
};

const AnnonceCard = (): any => {
  const [totalReactPackages, setTotalReactPackages] = useState<Property[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/home")
      .then(response => response.json())
      .then(data => {
        console.log(data);
        setTotalReactPackages(data);
      })
  }, []);

  /*
  const property = {
    imageUrl: totalReactPackages[0].imageUrl,
    NamePoste: "",
    entreprise: "apple",
    secteur: "informatique",
    contract: "CDI",
    localisation: "Paris",
    date: "13/04/2022",
    title: 'Dev Full stack',
    remuneration: '$100/D',
  }*/
  {

    return (
      totalReactPackages.map((data) =>
        <HStack w='100%' h="130px" borderWidth='2px' borderRadius='lg' px="16px" py="4px" justify="space-between">
          <HStack>
            <Image boxSize="100px" src={data.Image} borderRadius="full" />
            <VStack p='4' align="left">
              <HStack alignItems='baseline' spacing="8px">
                <Badge borderRadius='full' px='3' >
                  {data.Name}
                </Badge>
                <Badge borderRadius='full' px='2' colorScheme='teal'>
                  New
                </Badge>
              </HStack>
              <HStack alignItems='baseline' spacing="8px">
                <Text
                  as='h4'
                  fontWeight='semibold'
                  lineHeight='tight'
                  isTruncated
                >
                  {data.Contrat}
                </Text>
                <Badge colorScheme='green'>Verifier</Badge>
              </HStack>
              <HStack alignItems='baseline' spacing="8px">

                <Badge borderRadius='full' px='3' fontSize='18px' >
                  {data.Localisation}
                </Badge>
                <Badge borderRadius='full' px='3' fontSize='18px'>
                  {data.Remuneration}
                </Badge>
              </HStack>
            </VStack>
          </HStack>
          <Button colorScheme="blue" borderRadius="40">
            Plus d'information
          </Button>
        </HStack>
      )
    )
  }
}

export default AnnonceCard;