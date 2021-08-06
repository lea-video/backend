package model.core;

import model.utility.Title;

public interface Playable {
    Title getTitle();
    Playable getNext(int current);
}
