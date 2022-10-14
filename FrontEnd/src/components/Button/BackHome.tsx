import { Box, Link} from '@chakra-ui/react';

export default function BackHome() {
    return (
      <Box textAlign="center" py={1} mr={12}>
          <Link href="/home" >Home</Link>
      </Box>
    );
  }