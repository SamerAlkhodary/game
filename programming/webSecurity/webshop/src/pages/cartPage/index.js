import React, { useContext, useEffect, useState } from 'react';

import { Grid, List, ListItem, ListItemAvatar, Avatar, ListItemText, Typography, Divider, Container } from '@material-ui/core';
import useStyles from './style';
import CartContext from '../../contexts/cartContext';

const CartPage = () => {
    const classes = useStyles();
   
    const { cartState } = useContext(CartContext);
    const total = cartState.items.reduce(function (accumulator, item) {
        return accumulator + item.price;
      }, 0);
  
    const listItems = cartState.items.map(
        item =>
            <>
                <ListItem alignItems="flex-start">
                    <ListItemAvatar>
                        <Avatar alt="Remy Sharp" src={item.img} className={classes.large} />
                    </ListItemAvatar>
                    <ListItemText
                        primary={item.name}
                        secondary={
                            <React.Fragment>
                                <Typography
                                    component="span"
                                    variant="body2"
                                    className={classes.inline}
                                    color="textPrimary"
                                >
                                    {item.price + " " + item.unit}
                                </Typography>

                            </React.Fragment>
                        }
                    />
                </ListItem>
                <Divider variant="inset" component="li" />
            </>


    );
    return (
        <div className={classes.root}>
            <main>

                <Grid container justify="space-around">
                   <Grid item>
                   <List className={classes.list}>
                        {listItems}
                    </List>
                   </Grid>
                   <Grid item>
                   <Typography className={classes.totalAmount}>
                        {"Amount is "+ total + " SEK"}
                    </Typography>
                   </Grid>
                </Grid>

            </main>
        </div>
    );

}
export default CartPage;