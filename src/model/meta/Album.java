package model.meta;

import model.core.Playable;
import model.utility.Celebrity;
import model.utility.Title;

import java.util.ArrayList;
import java.util.List;

public class Album implements Playable {
    Title title;
    List<Song> songs;

    public Album(Title title) {
        this.title = title;
        this.songs = new ArrayList<>();
    }

    public void addSong(Song s) {
        this.songs.add(s);
    }

    @Override
    public Title getTitle() {
        return this.title;
    }

    @Override
    public Playable getNext(int current) {
        if (current + 1 >= this.songs.size()) {
            return null;
        }
        return songs.get(current + 1);
    }
}
