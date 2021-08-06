package model.user;

public interface Credentials {
    Session getSession();
    boolean isValid();
}
