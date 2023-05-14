import { Box, Button, Card, CardActions, CardContent, Grid, Typography } from "@mui/material";

const Signin = () => {
  return (
    <Box
      display="flex"
      alignItems="center"
      justifyContent="center"
      style={{ height: "100vh", width: "100vw" }}
    >
      <Card variant="outlined" sx={{ width: "300px", height: "300px" }}>
        <CardContent>
          <Box display="flex" alignItems="center" justifyContent="center">
            <Typography>sign in form</Typography>
          </Box>
        </CardContent>
        <CardActions>
          <Button variant="contained" sx={{ textTransform: "none" }}>
            <Typography>sign in</Typography>
          </Button>
        </CardActions>
      </Card>
    </Box>
  );
};

export default Signin;
