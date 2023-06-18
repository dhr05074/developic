import { useRecoilState } from "recoil";
import { profileState } from "../recoil/profile.recoil";
import { useEffect, useState } from "react";
import { api } from "@/api/defaultApi";
import { v4 as uuidv4 } from "uuid";
import { Record } from "api/api";
import { t } from "msw/lib/glossary-de6278a9";
import { useNavigate } from "react-router-dom";
import { loadingState } from "@/recoil/component.recoil";

function useProfile() {
    const navigate = useNavigate();

    const [profile, setProfile] = useRecoilState(profileState);
    const [records, setRecords] = useState<Record[]>();
    const [isLoading, setLoading] = useRecoilState(loadingState);
    const setAuth = async () => {
        if (!profile.headers.Authorization) {
            const uuid: string = uuidv4();
            // const name = btoa("JaeHwan"); // 대안.

            setProfile((prevState) => {
                return {
                    ...prevState,
                    headers: { Authorization: uuid, "Content-Type": "application/json" },
                    nickname: uuid,
                };
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
                            "Content-Type": "application/json",
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
    const getSingleRecord = (recordId: string) => {
        api.getRecord(recordId, { headers: profile.headers })
            .then((record) => {
                setLoading(false);
                navigate("/result");
                console.log("getRecord", record);
            })
            .catch((error) => {
                console.log("getSingleRecord : ", error.response.data);
                if (error.response.status === 409) {
                    // 아직 채점안됌.
                    setTimeout(() => {
                        getSingleRecord(recordId);
                    }, 3000);
                }
            });
    };
    return { profile, records, getProfile, getRecords, getSingleRecord, setAuth };
}

export default useProfile;
