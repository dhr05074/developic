import { useRecoilState } from "recoil";
import { profileState } from "../recoil/profile.recoil";
import { useEffect, useState } from "react";
import { api } from "@/api/defaultApi";
import { v4 as uuidv4 } from "uuid";
import { Record } from "api/api";
import { t } from "msw/lib/glossary-de6278a9";

function useProfile() {
    const [profile, setProfile] = useRecoilState(profileState);
    const [records, setRecords] = useState<Record[]>();
    const setAuth = async () => {
        if (!profile.nickname) {
            const uuid: string = uuidv4();
            // const name = btoa("JaeHwan"); // 대안.

            setProfile((prevState) => {
                return { ...prevState, headers: { Authorization: uuid }, nickname: uuid };
            });
            console.log("profile.hook", profile);
            return true;
        }
        return true;
    };
    const getProfile = () => {
        setAuth().then(() => {
            console.log("getProfile", profile);
            api.getMe({ headers: profile.headers })
                .then((me) => {
                    const getProfile = {
                        nickname: me.data.nickname,
                        elo_score: me.data.elo_score,
                        headers: {
                            Authorization: me.data.nickname,
                        },
                    };
                    setProfile(getProfile);
                })
                .catch((error) => {
                    console.log("getMe Error : ", error);
                });
        });
    };
    const getRecords = () => {
        api.getRecords(undefined, undefined, { headers: profile.headers })
            .then((records) => {
                console.log("🚀 ~ file: Profile.hook.tsx:35 ~ .then ~ records:", records);
                setRecords(records.data.records);
            })
            .catch((error) => {
                console.log("getRecords Error : ", error);
            });
        return () => {};
    };
    return { profile, records, getProfile, getRecords, setAuth };
}

export default useProfile;
