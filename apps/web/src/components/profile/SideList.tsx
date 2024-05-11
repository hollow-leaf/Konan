"use client";
import React, { useState } from "react";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import Collapse from "@mui/material/Collapse";
import ExpandLess from "@mui/icons-material/ExpandLess";
import ExpandMore from "@mui/icons-material/ExpandMore";
import Checkbox from "@mui/material/Checkbox";
import { bg_color, hat, glasses } from "@/content/SideBarObject";

export function NestedList() {
  const [openIndex, setOpenIndex] = useState<number | unknown>(null);
  const [check, setCheck] = useState<number | unknown>(null);
  const [isChecked, setIsChecked] = React.useState<boolean[]>(
    new Array(6).fill(false),
  );
  const handleClick = (index: number) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  const handleChecked = (id: number) => {
    // If the same item is clicked again, toggle its checked state
    if (check === id) {
      const updatedCheckedState = isChecked.map((checked, index) =>
        index === id ? !checked : checked,
      );
      setIsChecked(updatedCheckedState);
    } else {
      // If a different item is clicked, only set its checked state to true
      const updatedCheckedState = isChecked.map((checked, index) =>
        index === id ? true : false,
      );
      setIsChecked(updatedCheckedState);
    }
    setCheck(id);
  };

  return (
    <List sx={{ width: "100%", maxWidth: 360, bgcolor: "background.paper" }}>
      <ListItemButton onClick={() => handleClick(0)}>
        <ListItemText primary="Backround" />
        {openIndex === 0 ? <ExpandLess /> : <ExpandMore />}
      </ListItemButton>
      {bg_color.map((item, index) => (
        <Collapse key={index} in={openIndex === 0} timeout="auto" unmountOnExit>
          <List component="div" disablePadding>
            <ListItemButton
              sx={{ pl: 4 }}
              onClick={() => handleChecked(item.id)}
            >
              <Checkbox checked={isChecked[item.id]} />
              <ListItemText
                primary={`${item.color}`}
                secondary={`${item.amount}`}
              />
            </ListItemButton>
          </List>
        </Collapse>
      ))}

      <ListItemButton onClick={() => handleClick(1)}>
        <ListItemText primary="Hat" />
        {openIndex === 1 ? <ExpandLess /> : <ExpandMore />}
      </ListItemButton>
      {hat.map((item, index) => (
        <Collapse key={index} in={openIndex === 1} timeout="auto" unmountOnExit>
          <List component="div" disablePadding>
            <ListItemButton
              sx={{ pl: 4 }}
              onClick={() => handleChecked(item.id)}
            >
              <Checkbox checked={isChecked[item.id]} />
              <ListItemText
                primary={`${item.hat}`}
                secondary={`${item.amount}`}
              />
            </ListItemButton>
          </List>
        </Collapse>
      ))}

      <ListItemButton onClick={() => handleClick(2)}>
        <ListItemText primary="Glasses" />
        {openIndex === 2 ? <ExpandLess /> : <ExpandMore />}
      </ListItemButton>
      {glasses.map((item, index) => (
        <Collapse key={index} in={openIndex === 2} timeout="auto" unmountOnExit>
          <List component="div" disablePadding>
            <ListItemButton
              sx={{ pl: 4 }}
              onClick={() => handleChecked(item.id)}
            >
              <Checkbox checked={isChecked[item.id]} />
              <ListItemText
                primary={`${item.glasses}`}
                secondary={`${item.amount}`}
              />
            </ListItemButton>
          </List>
        </Collapse>
      ))}
    </List>
  );
}
